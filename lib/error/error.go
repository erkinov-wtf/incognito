package errorhandling

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func LogError(err error, message string) {
	log.Println(message+":", err)
}

func SendResponseToUser(bot *tgbotapi.BotAPI, message string, chatID int64) {
	response := tgbotapi.NewMessage(chatID, message)
	_, err := bot.Send(response)
	if err != nil {
		LogError(err, "Error sending response to user")
	}
}
