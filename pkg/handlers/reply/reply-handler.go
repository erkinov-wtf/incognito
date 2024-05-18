package reply

import "github.com/NicoNex/echotron/v3"

func HandleForwardedReply(api *echotron.API, message *echotron.Message, chatID int64) {
	// Forward the original message to the channel
	forwardOptions := &echotron.ForwardOptions{}
	api.ForwardMessage(chatID, message.ReplyToMessage.Chat.ID, message.ReplyToMessage.ID, forwardOptions)

	// Then send the user's reply
	api.SendMessage(message.Text, chatID, &echotron.MessageOptions{
		ReplyMarkup: nil, // No reply markup for now
	})
}
