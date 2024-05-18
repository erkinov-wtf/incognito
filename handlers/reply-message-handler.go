package handlers

import "github.com/NicoNex/echotron/v3"

func ReplyMessage(api *echotron.API, message *echotron.Message) {
	replyText := "Got your message!"
	options := &echotron.MessageOptions{
		ReplyMarkup: nil, // No reply markup for now
	}

	// Send the reply, referencing the original message
	api.SendMessage(replyText, message.Chat.ID, options)
}
