package main

import (
	"log"

	"github.com/beenson/URL_Shortener/pkg/migrate"
	"github.com/beenson/URL_Shortener/pkg/repository"
	route "github.com/beenson/URL_Shortener/pkg/routes"
	util "github.com/beenson/URL_Shortener/pkg/utils"
	"github.com/beenson/URL_Shortener/service/cache"
	"github.com/beenson/URL_Shortener/service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// @title       URL Shortener API
// @version     1.0
// @description URL Shortener API
// @BasePath    /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init
	util.Init()
	repository.Init()

	// database
	database.DbInit()
	migrate.Migrate()

	// cache
	cache.Init()

	// routes
	app := fiber.New()
	route.PublicRoutes(app)
	route.SwaggerRoute(app)
	route.NotFoundRoute(app)

	// Listen port 80
	app.Listen(":80")
}
