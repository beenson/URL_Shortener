package util

import (
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

func Init() {
	// code_generate
	rand.Seed(time.Now().Unix())

	RandFunc = func(max int) int {
		return rand.Intn(max)
	}

	// parse
	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}
