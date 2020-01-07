package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

func info(command string, telegramId int) (response string) {
	response = "Бот контроля личных финансовых. Попробуйте отправить \"1000  Products\" "
	return response
}

func newTransaction(command string, telegramId int) (response string) {
	namePattern := regexp.MustCompile(`\s(.*)`)
	name := strings.TrimSpace(namePattern.FindString(command))

	amountPattern := regexp.MustCompile(`(\d+)`)
	amountString := amountPattern.FindString(command)
	amount, err := strconv.Atoi(amountString)

	if err != nil || amount == 0 {
		return
	}

	addTransaction(telegramId, amount, name)
	response = "Добавлено. Остаток сегодня: num"
	return response
}

func Today(command string, telegramId int) (response string) {
	var transactions []Transaction
	db.Raw("SELECT*FROM transactions WHERE user_id = ? and DATE(date) = ?", userId(telegramId), time.Now().Format("2006-01-02")).Scan(&transactions)
	if transactions == nil {
		response = "Сегодня нет трат"
		return
	}
	response = "Траты сегодня:\n"
	for _, transaction := range transactions {
		response += transaction.Name + " - " + strconv.Itoa(transaction.Amount) + "\n"
	}
	return
}

type Budget struct {
	Income      int
	TodayBudget int
}

func GetBudget(command string, telegramId int) (response string) {
	var budget Budget
	db.Table("transactions").Select("sum(amount) as income").
		Where(map[string]interface{}{"type": "month", "user_id": userId(telegramId)}).
		Scan(&budget)
	db.Table("transactions").Select("sum(amount) as income").
		Where(map[string]interface{}{"type": "month", "user_id": userId(telegramId)}).
		Scan(&budget)

	response = strconv.Itoa(budget.Income)
	return
}
