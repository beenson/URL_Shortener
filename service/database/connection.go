package database

import (
	"fmt"
	"os"
	"time"

	// V2需要引用這package
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func DbInit() {
	// 資料庫參數設定
	var (
		UserName     string = os.Getenv("DB_USERNAME")
		Password     string = os.Getenv("DB_PASSWORD")
		Network      string = os.Getenv("DB_NETWORK")
		Addr         string = os.Getenv("DB_HOST")
		Port         string = os.Getenv("DB_PORT")
		Database     string = os.Getenv("DB_DATABASE")
		MaxLifetime  int    = 10
		MaxOpenConns int    = 10
		MaxIdleConns int    = 10
	)

	addr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8&parseTime=True", UserName, Password, Network, Addr, Port, Database)
	// 連接MySQL
	conn, err := gorm.Open(mysql.Open(addr), &gorm.Config{})
	if err != nil {
		fmt.Println("🔴 Connection to MySQL failed:", err)
		return
	}
	// 設定ConnMaxLifetime/MaxIdleConns/MaxOpenConns
	db, err1 := conn.DB()
	if err1 != nil {
		fmt.Println("🔴 Get Database failed:", err)
		return
	}
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	db.SetMaxIdleConns(MaxIdleConns)
	db.SetMaxOpenConns(MaxOpenConns)
	Instance = conn
	fmt.Println("🟢 DB Connection Init Success.")
}
