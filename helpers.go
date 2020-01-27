package main

import (
	"flag"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/joho/godotenv"
	"regexp"
)

type Config map[string]string

func GetConfig() Config {
	configFile := flag.String("env", ".env", "environment file path")
	flag.Parse()
	config, _ := godotenv.Read(*configFile)
	return config
}

func userId(tgId int) uint {
	var user User
	db.FirstOrCreate(&user, User{TelegramId: tgId})
	return user.ID
}

type Router struct {
	routes []*Route
}

type Route struct {
	Handler   func(string, int) string
	Message   string
	IsPattern bool
}

func (r *Router) Add(message string, handler func(string, int) string, isPattern bool) {
	route := &Route{
		handler,
		message,
		isPattern,
	}
	r.routes = append(r.routes, route)
}

func (r *Router) DoCommand(message string, telegramId int) (response string) {
	for _, route := range r.routes {
		if route.IsPattern == true {
			isMatch, _ := regexp.MatchString(route.Message, message)
			if isMatch == true {
				response = route.Handler(message, telegramId)
				return
			}
		}
		if route.Message == message {
			response = route.Handler(message, telegramId)
			return
		}
	}
	return
}

func initBot() {
	var r Router
	r.Get()
	bot, err := tgbotapi.NewBotAPI(config["TELEGRAM_TOKEN"])
	if err != nil {
		panic(err)
	}
	fmt.Println("Authorized on account %s", bot.Self.UserName)

	_, err = bot.RemoveWebhook()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 5
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Printf("[%s] %s %s - %s \n", update.Message.From.UserName, update.Message.From.FirstName, update.Message.From.LastName, update.Message.Text)
		reply := r.DoCommand(update.Message.Text, update.Message.From.ID)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}
