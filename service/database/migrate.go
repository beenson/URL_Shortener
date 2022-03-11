package database

import (
	model "github.com/beenson/URL_Shortener/app/models"
)

func Migrate() {
	Instance.Migrator().AutoMigrate(&model.Shorten{})
}
