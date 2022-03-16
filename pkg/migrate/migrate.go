package migrate

import (
	model "github.com/beenson/URL_Shortener/app/models"
	"github.com/beenson/URL_Shortener/service/database"
)

func Migrate() {
	database.Instance.Migrator().AutoMigrate(&model.Shorten{})
}
