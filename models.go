package main

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model
	ID         int
	TelegramId int
}

type Transaction struct {
	gorm.Model
	UserID int
	User   User
	Date   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Type   string
	Name   string
	Amount float64
}

var db *gorm.DB //база данных

func init() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=fman password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{}, &Transaction{})
}

func GetDB() *gorm.DB {
	return db
}
