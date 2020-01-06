package main

func (r *Router) Get() {
	r.Add("/start", info, false)
	r.Add(`\d\s\w`, newTransaction, true)
}
