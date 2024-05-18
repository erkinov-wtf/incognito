package message

import (
	"github.com/NicoNex/echotron/v3"
	"pu/handlers"
	"pu/pkg/handlers/media"
	"pu/pkg/handlers/reply"
	"pu/pkg/types"
	"sync"
)

func HandleMessage(api *echotron.API, message *echotron.Message, chatID int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, mu *sync.Mutex) {
	switch {
	case message.Text != "":
		// If the message is a text, reply to it
		api.SendMessage(message.Text, chatID, nil)
	case message.Photo != nil:
		// If the message is a photo, handle media group
		media.HandleMediaGroup(api, message, chatID, mediaGroupMap, mu, types.Photo)
	case message.Document != nil:
		// If the message is a document, handle media group
		media.HandleMediaGroup(api, message, chatID, mediaGroupMap, mu, types.Doc)
	}

	// If the message is a reply to another message, handle the reply
	if message.ReplyToMessage != nil {
		reply.HandleForwardedReply(api, message, chatID)
	}

	// Reply to the user
	handlers.ReplyMessage(api, message)
}
