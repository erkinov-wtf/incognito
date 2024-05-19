package handlers

import (
	"github.com/NicoNex/echotron/v3"
	"pu/pkg/replies/messages"
)

func ReplyMessage(api *echotron.API, message *echotron.Message) {
	replyText := messages.GetRandomReply()
	options := &echotron.MessageOptions{
		ReplyMarkup: nil,
		ParseMode:   echotron.Markdown,
	}

	// Send the replies, referencing the original message
	api.SendMessage(replyText, message.Chat.ID, options)
}
