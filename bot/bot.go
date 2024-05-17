package bot

import (
	"log"
	"pu/message"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(botToken string, channelUsername string, chatID int64) {
	if botToken == "" || channelUsername == "" {
		log.Fatal("Please set the TELEGRAM_BOT_TOKEN and USERNAME environment variables")
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	var mu sync.Mutex
	var queue []*tgbotapi.Message

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		mu.Lock()
		queue = append(queue, update.Message)
		mu.Unlock()
		go message.ProcessMessage(bot, &mu, &queue, channelUsername, chatID)
	}
}
