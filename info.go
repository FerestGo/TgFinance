package main

import (
	"regexp"
	"strconv"
)

func info(command string, telegramId int) (response string) {
	response = "Бот контроля личных финансовых. Попробуйте отправить \"1000  Products\" "
	return response
}

func newTransaction(command string, telegramId int) (response string) {
	namePattern := regexp.MustCompile(`^\S*\s+(\S+)`)
	name := namePattern.FindString(command)
	// todo fix it
	amountPattern := regexp.MustCompile(`(\d+)`)
	amountString := amountPattern.FindString(command)
	amount, err := strconv.ParseFloat(amountString, 64)
	if err != nil || amount == 0 {
		return
	}
	addTransaction(telegramId, amount, name)
	response = "Добавлено. Остаток сегодня: num"
	return response
}
