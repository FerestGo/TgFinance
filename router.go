package main

func (r *Router) Get() {
	r.Add("/start", info, false)
	r.Add(`^\d+\s.*$`, newTransaction, true)
	r.Add(`/today`, Today, false)
	r.Add(`Сегодня`, Today, false)
	r.Add(`/budget`, GetBudget, false)
	r.Add(`Бюджет`, GetBudget, false)
	r.Add(`Месяц`, Month, false)
	r.Add(`/month`, Month, false)
	r.Add(`Регулярный`, AddIncome, false)
}
