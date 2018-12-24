package main

import (
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
)

func main() {

	//fmt.Println(configuration.TelegramBotToken)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("783917145:AAHOcCXfORlcKaJ1leKttbs9apz1baIsZVA"))

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}
	// В канал updates будут приходить все новые сообщения.
	for update := range updates {
		// Создав структуру - можно её отправить обратно боту
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}
