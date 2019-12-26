package main

import (
	"fmt"
	"regexp"
)

func main() {
	db := GetDB()
	var reply string
	pattern := `\d\s\w`
	command := "1000 Products"
	flag, _ := regexp.MatchString(pattern, command)

	if command == "/start" {
		reply = "Это телега бот"
	}
	if flag == true {
		var transaction Transaction
		db.Last(&transaction)
		fmt.Println(&transaction)
		fmt.Println(reply)
	}
	fmt.Println("thats all")
	// db := InitDb()
	// Чтение
	// user := new(User)
	// transaction := new(Transaction)

	// var user User
	// var transaction Transaction
	// db.First(&user, 1) // find product with id 1
	// addTranstction(11, 11, "Products")
	// db.Last(&transaction)
	// fmt.Println(&transaction)

}

// func addTranstction(TelegramId int, Amount float64, Name string) {
//  db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=fman password=root sslmode=disable")
//  if err != nil {
//   panic(err)
// }
// defer db.Close()
// var user User
// // var transaction Transaction
// db.Where("telegram_id = ?", TelegramId).First(&user)
// db.Create(&Transaction{
//   UserID:user.ID,
//   Type:"day",
//   Name:Name,
//   Amount: Amount})

// }
