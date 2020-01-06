package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	ID         uint `gorm:"primary_key"`
	TelegramId int
}

type Transaction struct {
	ID     uint `gorm:"primary_key"`
	UserID uint
	User   User
	Date   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Type   string
	Name   string
	Amount float64
}

var db *gorm.DB

func initDB() *gorm.DB {
	config = GetConfig() // todo get config in main()
	agrs := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config["POSTGRES_HOST"], config["POSTGRES_PORT"], config["DATABASE_USER"], config["DATABASE_NAME"], config["DATABASE_PASSWORD"])
	db, err := gorm.Open("postgres", agrs)
	if err != nil {
		panic(err)
	}
	//db.LogMode(true)
	db.AutoMigrate(&User{}, &Transaction{})
	return db
}
