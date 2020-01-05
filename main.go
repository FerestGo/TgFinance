package main

import (
	"fmt"
	"regexp"
)

var config Config

func main() {
	db = initDB()
	var reply string
	pattern := `\d\s\w`
	command := "1000 Products"
	_, _ = regexp.MatchString(pattern, command)
	fmt.Print(reply)

	if command == "/start" {
		reply = "Это телега бот"
	}

	//if flag == true {
	//	db.Create(&User{
	//		TelegramId: 1})
	//	var user User
	//	db.Last(&user)
	//	fmt.Println(&user)
	//}
	addTransaction(1, 11, "Products")
	addTransaction(1, 777, "Taxi")

}

func addTransaction(TelegramId int, Amount float64, Name string) {
	var user User
	db.Where("telegram_id = ?", TelegramId).First(&user)
	db.Create(&Transaction{
		UserID: user.ID,
		Type:   "day",
		Name:   Name,
		Amount: Amount})
}
