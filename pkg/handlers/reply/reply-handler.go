package reply

import (
	"sync"

	"github.com/NicoNex/echotron/v3"
)

func HandleReplyMessage(api *echotron.API, message *echotron.Message, chatID int64, originalMessages *map[int64]int, mu *sync.Mutex) {
	// Forward the original message to the channel
	forwardOptions := &echotron.ForwardOptions{}
	api.ForwardMessage(chatID, message.ReplyToMessage.Chat.ID, message.ReplyToMessage.ID, forwardOptions)

	// Then send the user's replies
	api.SendMessage(message.Text, chatID, &echotron.MessageOptions{
		ReplyMarkup: nil, // No replies markup for now
	})
}
