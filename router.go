package main

func (r *Router) Get() {
	r.Add("/start", start, false)
	r.Add("/help", info, false)
	r.Add("/info", info, false)

	r.Add(`^\d+\s.*$`, newTransaction, true)
	r.Add(`^.*\s\d+$`, newTransaction, true)

	r.Add(`/today`, Today, false)
	r.Add(`Сегодня`, Today, false)

	r.Add(`/budget`, GetBudget, false)
	r.Add(`Бюджет`, GetBudget, false)
	r.Add(`Б`, GetBudget, false)

	r.Add(`Месяц`, Month, false)
	r.Add(`/month`, Month, false)

	r.Add(`^\+\+\d+\s.*$`, AddRegularIncome, true)
	r.Add(`^\+\+\s\d+\s.*$`, AddRegularIncome, true)

	r.Add(`^Регулярный\s\d+\s.*$`, AddRegularCost, true)
	r.Add(`^--\d+\s.*$`, AddRegularCost, true)
}
