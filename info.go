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
	amount = amount * -1

	if err != nil || amount == 0 {
		return
	}

	addTransaction(telegramId, amount, name)
	var budget Budget
	budget.Get(userId(telegramId))
	response = "Добавлено. Остаток сегодня: " + strconv.Itoa(budget.TodayBudget)
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
	DailyBudget int
	SumCosts    int
	TodayBudget int
}

func (budget *Budget) Get(userId uint) {
	db.Table("transactions").Select("sum(amount) as income").
		Where(map[string]interface{}{"type": "month", "user_id": userId}).
		Scan(budget)
	budget.DailyBudget = budget.Income / 30 // todo change it
	db.Table("transactions").Select("sum(amount) as sum_costs").
		Where(map[string]interface{}{"type": "day", "user_id": userId}).
		Scan(budget)
	currentDay, _ := strconv.Atoi(time.Now().Format("02"))
	budget.TodayBudget = (budget.DailyBudget * currentDay) + budget.SumCosts
}

func GetBudget(command string, telegramId int) (response string) {
	var budget Budget
	budget.Get(userId(telegramId))

	response += "Прибыль в месяц: " + strconv.Itoa(budget.Income) + "\n"
	response += "Расходы в этом месяце: " + strconv.Itoa(budget.SumCosts) + "\n"
	response += "Ежедневный бюджет: " + strconv.Itoa(budget.DailyBudget) + "\n"
	response += "Остаток на сегодня: " + strconv.Itoa(budget.TodayBudget)
	return
}

func Month(command string, telegramId int) (response string) {
	var transactions []Transaction
	db.Raw("SELECT*FROM transactions WHERE user_id = ? and date_part('month', date) = ? and type = 'day'", userId(telegramId), time.Now().Format("01")).Scan(&transactions)
	if transactions == nil {
		response = "В этом месяце нет трат"
		return
	}
	response = "В этом месяце:\n"
	for _, transaction := range transactions {
		response += "[" + transaction.Date.Format("02") + "] " + transaction.Name + " " + strconv.Itoa(transaction.Amount) + "\n"
	}
	return
}

func AddRegularIncome(command string, telegramId int) (response string) {
	namePattern := regexp.MustCompile(`\s(.*)`)
	name := strings.TrimSpace(namePattern.FindString(command))

	amountPattern := regexp.MustCompile(`(\d+)`)
	amountString := amountPattern.FindString(command)
	amount, err := strconv.Atoi(amountString)

	if err != nil || amount == 0 {
		return
	}

	db.Create(&Transaction{
		UserID: userId(telegramId),
		Type:   "month",
		Name:   name,
		Amount: amount})
	response = "Добавлен регулярный доход: " + strconv.Itoa(amount)
	return response
}

func AddRegularCost(command string, telegramId int) (response string) {
	namePattern := regexp.MustCompile(`\s(.*)`)
	name := strings.TrimSpace(namePattern.FindString(command))

	amountPattern := regexp.MustCompile(`(\d+)`)
	amountString := amountPattern.FindString(command)
	amount, err := strconv.Atoi(amountString)
	amount = amount*-1

	if err != nil || amount == 0 {
		return
	}

	db.Create(&Transaction{
		UserID: userId(telegramId),
		Type:   "month",
		Name:   name,
		Amount: amount})
	response = "Добавлен регулярный расход: " + strconv.Itoa(amount)
	return response
}