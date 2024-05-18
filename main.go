package main

import (
	"sync"
	"time"

	"github.com/NicoNex/echotron/v3"
	"pu/internal/config"
)

func main() {
	cfg := config.MustLoad()
	api := echotron.NewAPI(cfg.Token)

	mediaGroupMap := make(map[string][]echotron.GroupableInputMedia)
	var mu sync.Mutex

	// Listen for messages
	for u := range echotron.PollingUpdates(cfg.Token) {
		if u.Message != nil {
			handleMessage(&api, u.Message, cfg.ChannelID, &mediaGroupMap, &mu)
		}
	}
}

func handleMessage(api *echotron.API, message *echotron.Message, chatID int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, mu *sync.Mutex) {
	switch {
	case message.Text != "":
		// If the message is a text, reply to it
		api.SendMessage(message.Text, chatID, nil)
	case message.Photo != nil:
		// If the message is a photo, handle media group
		handleMediaGroup(api, message, chatID, mediaGroupMap, mu, "photo")
	case message.Document != nil:
		// If the message is a document, handle media group
		handleMediaGroup(api, message, chatID, mediaGroupMap, mu, "document")
	}

	// If the message is a reply to another message, handle the reply
	if message.ReplyToMessage != nil {
		handleForwardedReply(api, message, chatID)
	}

	// Reply to the user
	replyMessage(api, message)
}

func handleMediaGroup(api *echotron.API, message *echotron.Message, chatID int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, mu *sync.Mutex, mediaType string) {
	var media echotron.GroupableInputMedia

	if mediaType == "photo" {
		media = echotron.InputMediaPhoto{
			Type:    "photo",
			Media:   echotron.NewInputFileID(message.Photo[len(message.Photo)-1].FileID),
			Caption: message.Caption,
		}
	} else if mediaType == "document" {
		media = echotron.InputMediaDocument{
			Type:    "document",
			Media:   echotron.NewInputFileID(message.Document.FileID),
			Caption: message.Caption,
		}
	}

	if message.MediaGroupID != "" {
		// Use mutex to safely access the mediaGroupMap
		mu.Lock()
		(*mediaGroupMap)[message.MediaGroupID] = append((*mediaGroupMap)[message.MediaGroupID], media)
		mu.Unlock()

		// Send media group after a delay to ensure all items are collected
		go func(mediaGroupID string) {
			time.Sleep(2 * time.Second) // Adjust delay as necessary
			mu.Lock()
			if mediaGroup, ok := (*mediaGroupMap)[mediaGroupID]; ok {
				if len(mediaGroup) > 1 {
					api.SendMediaGroup(chatID, mediaGroup, nil)
					delete(*mediaGroupMap, mediaGroupID)
				}
			}
			mu.Unlock()
		}(message.MediaGroupID)
	} else {
		// If it's a single media, send it directly
		if mediaType == "photo" {
			photoID := message.Photo[len(message.Photo)-1].FileID
			api.SendPhoto(echotron.NewInputFileID(photoID), chatID, &echotron.PhotoOptions{
				Caption: message.Caption,
			})
		} else if mediaType == "document" {
			fileID := message.Document.FileID
			api.SendDocument(echotron.NewInputFileID(fileID), chatID, &echotron.DocumentOptions{
				Caption: message.Caption,
			})
		}
	}
}

func handleForwardedReply(api *echotron.API, message *echotron.Message, chatID int64) {
	// Forward the original message to the channel
	forwardOptions := &echotron.ForwardOptions{}
	api.ForwardMessage(chatID, message.ReplyToMessage.Chat.ID, message.ReplyToMessage.ID, forwardOptions)

	// Then send the user's reply
	api.SendMessage(message.Text, chatID, &echotron.MessageOptions{
		ReplyMarkup: nil, // No reply markup for now
	})
}

func replyMessage(api *echotron.API, message *echotron.Message) {
	replyText := "Got your message!"
	options := &echotron.MessageOptions{
		ReplyMarkup: nil, // No reply markup for now
	}

	// Send the reply, referencing the original message
	api.SendMessage(replyText, message.Chat.ID, options)
}
