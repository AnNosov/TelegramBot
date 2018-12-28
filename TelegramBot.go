package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"os"
	"time"
)

type Config struct {
	TelegramBotToken string
}

func main() {
	file, _ := os.Open("config.json") // зашил токен в json
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(configuration.TelegramBotToken)

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// u - структура с конфигом для получения апдейтов
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	// используя конфиг u создаем канал в который будут прилетать новые сообщения
	updates, err := bot.GetUpdatesChan(u)

	// в канал updates прилетают структуры типа Update
	// вычитываем их и обрабатываем
	for update := range updates {
		var reply string
		var a string
		var b string
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text) //логирование

		switch update.Message.Text { // обработка текста Command для команд :))
		case "/start":
			reply = "What?"
		case "Привет":
			reply = "Привет. Я телеграм-бот"
		case "Сколько времени?":
			t := time.Now()
			reply = t.Format(time.RFC822)
		case "Назови себя":
			reply = bot.Self.UserName
		case "Как меня зовут?":
			reply = update.Message.From.UserName
		case "Проверка":
			a = update.Message.From.FirstName
			b = update.Message.From.LastName
			reply = a + " " + b
			//default:
			//	reply = "I don't know that command" // на всякий
		}

		// создаем ответное сообщение
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, reply)
		// отправляем
		bot.Send(msg)
	}
}
