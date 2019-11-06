package main

import (
	"flag"
	"fmt"
	"github.com/Syfaro/telegram-bot-api"
	"log"
	"os"
)

var (
	telegramBotToken string
	config           Config
)

func init() {
	flag.StringVar(&telegramBotToken, "telegrambottoken", "757997705:AAGVgtuQJh2z4aK-Qb72WDpuw7-asIY-feM", "Telegram Bot Token")
	flag.Parse()
	if telegramBotToken == "" {
		log.Print("-telegrambottoken is required")
		os.Exit(1)
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config["TOKEN"])
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	_, err = bot.RemoveWebhook()
	fmt.Println(err)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		// универсальный ответ на любое сообщение
		reply := "Не знаю что сказать"
		if update.Message == nil {
			continue
		}

		// логируем от кого какое сообщение пришло
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		switch update.Message.Command() {
		case "start":
			reply = "Привет. Я телеграм-бот"
		case "hello":
			reply = "world"
		case "key":
			var markup tgbotapi.InlineKeyboardMarkup
			edit := tgbotapi.NewEditMessageText(
				update.Message.Chat.ID,
				update.Message.MessageID,
				"sample text",
			)
			edit.ReplyMarkup = &markup
			reply = "Привет. Я телеграм-бот"
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		bot.Send(msg)
	}
}
