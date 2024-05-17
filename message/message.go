package message

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	errorhandling "pu/lib/error"
	"sync"
)

func ProcessMessage(bot *tgbotapi.BotAPI, mu *sync.Mutex, queue *[]*tgbotapi.Message, channelUsername string, channelChatId int64) {
	mu.Lock()
	defer mu.Unlock()

	for len(*queue) > 0 {
		message := (*queue)[0]
		*queue = (*queue)[1:]

		switch {
		case message.Text != "":
			sendMessage(bot, message.Text, channelUsername, message.Chat.ID)

		case len(message.Photo) > 0:
			sendPhoto(bot, message.Photo, channelUsername, message.Chat.ID)

		case message.Document != nil:
			sendDocument(bot, message.Document, channelChatId, message.Chat.ID)

		default:
			sendMessage(bot, AllowedFileTypes, "", message.Chat.ID)
		}
	}
}

func sendMessage(bot *tgbotapi.BotAPI, message, channelUsername string, chatID int64) {
	if message != "/start" {
		msg := tgbotapi.NewMessageToChannel(channelUsername, message)
		_, err := bot.Send(msg)
		if err != nil {
			errorhandling.LogError(err, "Error sending message to channel")
			return
		}
		errorhandling.SendResponseToUser(bot, MessageToUser, chatID)
	}
}

func sendPhoto(bot *tgbotapi.BotAPI, photo []tgbotapi.PhotoSize, channelUsername string, chatID int64) {
	photoMsg := tgbotapi.NewPhotoToChannel(channelUsername, tgbotapi.FileID(photo[len(photo)-1].FileID))
	_, err := bot.Send(photoMsg)
	if err != nil {
		errorhandling.LogError(err, "Error sending photo to channel")
		return
	}
	errorhandling.SendResponseToUser(bot, PhotoToUser, chatID)
}

func sendDocument(bot *tgbotapi.BotAPI, doc *tgbotapi.Document, channelChatId int64, chatID int64) {
	if doc.FileSize > MaxFileSize {
		sendMessage(bot, FileSizeExceeds, "", chatID)
		return
	}
	fileID := tgbotapi.FileID(doc.FileID)
	docMsg := tgbotapi.NewDocument(channelChatId, fileID)
	_, err := bot.Send(docMsg)
	if err != nil {
		errorhandling.LogError(err, "Error sending document to channel")
		return
	}
	errorhandling.SendResponseToUser(bot, DocumentToUser, chatID)
}
