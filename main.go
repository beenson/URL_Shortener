package main

import (
	"log"

	"github.com/beenson/URL_Shortener/service/database"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 資料庫設定
	database.DbInit()
	database.Migrate()
}
