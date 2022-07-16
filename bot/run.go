package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	log "github.com/sirupsen/logrus"
	"os"
)

func Run(done chan<- bool, messages <-chan string) {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_SECRET"))
	if err != nil {
		log.Error("bot creation error", err)
		done <- true
		return
	}

	for {
		msg := tgbotapi.NewMessage(-675937582, <-messages)

		if _, err = bot.Send(msg); err != nil {
			log.Error("send message error", err)
		}
	}
}
