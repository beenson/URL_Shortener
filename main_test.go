package main

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	model "github.com/beenson/URL_Shortener/app/models"
	"github.com/beenson/URL_Shortener/pkg/migrate"
	"github.com/beenson/URL_Shortener/pkg/repository"
	route "github.com/beenson/URL_Shortener/pkg/routes"
	util "github.com/beenson/URL_Shortener/pkg/utils"
	"github.com/beenson/URL_Shortener/service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

type APItest struct {
	description  string // description of the test case
	method       string
	route        string
	requestBody  string
	expectedCode int
	expectedBody string
}

var app *fiber.App

func TestMain(m *testing.M) {
	// Setup
	setup()

	// Run test
	result := m.Run()

	// Teardown
	teardown()

	os.Exit(result)
}

func TestCreateAndRedirect(t *testing.T) {
	// Id not exist redirect failed
	testAPI(t, &APItest{
		description:  "get HTTP status 404, when /{id} not exists",
		method:       "GET",
		route:        "/abced",
		expectedCode: 404,
		expectedBody: `Not Found`,
	})

	// Create
	expireAt := time.Now().UTC().Add(time.Second * time.Duration(2))
	testAPI(t, &APItest{
		description:  "get HTTP status 200, when post /api/v1/urls with correct parameters",
		method:       "POST",
		route:        "/api/v1/urls",
		requestBody:  `{"url":"http://www.google.com/", "expireAt":"` + expireAt.Format(time.RFC3339) + `"}`,
		expectedCode: 200,
		expectedBody: `{"id":"abcde","shortUrl":"http://localhost/abcde"}`,
	})

	// Redirect
	testAPI(t, &APItest{
		description:  "get HTTP status 302, when /{id} exists",
		method:       "GET",
		route:        "/abcde",
		expectedCode: 302,
	})

	// Create another shorten url(reset RandFunc)
	resetRandomFunc()
	testAPI(t, &APItest{
		description:  "get HTTP status 200, when post /api/v1/urls with correct parameters",
		method:       "POST",
		route:        "/api/v1/urls",
		requestBody:  `{"url":"http://www.google.com/", "expireAt":"` + expireAt.Format(time.RFC3339) + `"}`,
		expectedCode: 200,
		expectedBody: `{"id":"fghij","shortUrl":"http://localhost/fghij"}`,
	})

	// Wait for expire
	for expireAt.After(time.Now()) {
		time.Sleep(100 * time.Millisecond)
	}
	// Expired
	testAPI(t, &APItest{
		description:  "get HTTP status 404, when /{id} expired",
		method:       "GET",
		route:        "/abced",
		expectedCode: 404,
		expectedBody: `Not Found`,
	})

	// Create again
	resetRandomFunc()
	expireAt = time.Now().UTC().Add(time.Second * time.Duration(2))
	testAPI(t, &APItest{
		description:  "get HTTP status 200, when post /api/v1/urls with correct parameters",
		method:       "POST",
		route:        "/api/v1/urls",
		requestBody:  `{"url":"http://www.google.com/", "expireAt":"` + expireAt.Format(time.RFC3339) + `"}`,
		expectedCode: 200,
		expectedBody: `{"id":"abcde","shortUrl":"http://localhost/abcde"}`,
	})
}

func TestCreateFail(t *testing.T) {
	tests := []APItest{
		{
			description:  "get HTTP status 400, when post /api/v1/urls without any parameters",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{}`,
			expectedCode: 400,
			expectedBody: `{"message":"url is required"}`,
		},
		{
			description:  "get HTTP status 400, when post /api/v1/urls without url",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{"expireAt":"2022-03-20T08:20:41Z"}`,
			expectedCode: 400,
			expectedBody: `{"message":"url is required"}`,
		},
		{
			description:  "get HTTP status 400, when post /api/v1/urls without expireAt",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{"url":"http://www.google.com/"}`,
			expectedCode: 400,
			expectedBody: `{"message":"expireAt is required"}`,
		},
		{
			description:  "get HTTP status 400, when post /api/v1/urls parameter url have wrong format",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{"url":"hello world", "expireAt":"2022-03-20T08:20:41Z"}`,
			expectedCode: 400,
			expectedBody: `{"message":"url should be url"}`,
		},
		{
			description:  "get HTTP status 400, when post /api/v1/urls parameter expireAt have wrong format",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{"url":"http://www.google.com/", "expireAt":"2022-03-2008:20:41Z"}`,
			expectedCode: 400,
			expectedBody: `{"message":"please check expireAt format"}`,
		},
		{
			description:  "get HTTP status 400, when post /api/v1/urls parameter expireAt already passed",
			method:       "POST",
			route:        "/api/v1/urls",
			requestBody:  `{"url":"http://www.google.com/", "expireAt":"2021-03-20T08:20:41Z"}`,
			expectedCode: 400,
			expectedBody: `{"message":"expire time has passed"}`,
		},
	}

	testAPIs(t, &tests)
}

func TestNotFound(t *testing.T) {
	tests := []APItest{
		{
			description:  "get HTTP status 404, when route is not exists",
			method:       "GET",
			route:        "/api/v1/url",
			expectedCode: 404,
			expectedBody: `Not Found`,
		},
		{
			description:  "get HTTP status 404, when method of route is not exists",
			method:       "GET",
			route:        "/api/v1/urls",
			expectedCode: 404,
			expectedBody: `Not Found`,
		},
	}

	testAPIs(t, &tests)
}

func setup() {
	if err := godotenv.Load("./.env.test"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Database Prepare
	database.DbInit()
	migrate.Migrate()

	// Init
	repository.Init()

	// routes
	app = fiber.New()
	route.PublicRoutes(app)
	route.NotFoundRoute(app)

	resetRandomFunc()
}

func resetRandomFunc() {
	var count int = -1
	util.RandFunc = func(max int) int {
		count += 1
		return count % max
	}
}

func teardown() {
	// Drop table
	database.Instance.Migrator().DropTable(&model.Shorten{})
}

func testAPIs(t *testing.T, tests *[]APItest) {
	for _, test := range *tests {
		testAPI(t, &test)
	}
}

func testAPI(t *testing.T, test *APItest) {
	req := httptest.NewRequest(test.method, test.route, strings.NewReader(test.requestBody))
	req.Header.Set("Content-Type", "application/json")

	resp, _ := app.Test(req, -1)

	assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

	buf := new(strings.Builder)
	io.Copy(buf, resp.Body)
	fmt.Println(buf.String())
	assert.Equalf(t, test.expectedBody, buf.String(), test.description)
}
