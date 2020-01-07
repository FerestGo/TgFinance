package main

var config Config

func main() {
	db = initDB()
	initBot()
}

func addTransaction(TelegramId int, Amount int, Name string) {
	db.Create(&Transaction{
		UserID: userId(TelegramId),
		Type:   "day",
		Name:   Name,
		Amount: Amount})
}
