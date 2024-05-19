package feedback

import (
	"github.com/NicoNex/echotron/v3"
	"pu/pkg/types"
	"sync"
	"time"
)

func SendFeedback(api *echotron.API, message *echotron.Message, feedbackChannelId int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()

	if message.MediaGroupID != "" {
		// No need for mediaType variable in media group handling
		var media echotron.GroupableInputMedia

		if message.Photo != nil {
			media = echotron.InputMediaPhoto{
				Type:    types.Photo,
				Media:   echotron.NewInputFileID(message.Photo[len(message.Photo)-1].FileID),
				Caption: message.Caption,
			}
		} else if message.Document != nil {
			media = echotron.InputMediaDocument{
				Type:    types.Doc,
				Media:   echotron.NewInputFileID(message.Document.FileID),
				Caption: message.Caption,
			}
		}

		(*mediaGroupMap)[message.MediaGroupID] = append((*mediaGroupMap)[message.MediaGroupID], media)

		// Use a goroutine to send the media group after a short delay to collect all media
		go func(mediaGroupID string) {
			time.Sleep(2 * time.Second) // Adjust delay as necessary
			mu.Lock()
			defer mu.Unlock()
			if mediaGroup, ok := (*mediaGroupMap)[mediaGroupID]; ok {
				if len(mediaGroup) > 1 {
					api.SendMediaGroup(feedbackChannelId, mediaGroup, nil)
					delete(*mediaGroupMap, mediaGroupID)
				}
			}
		}(message.MediaGroupID)
	} else {
		// Handle single media
		if message.Photo != nil {
			photoID := message.Photo[len(message.Photo)-1].FileID
			api.SendPhoto(echotron.NewInputFileID(photoID), feedbackChannelId, &echotron.PhotoOptions{
				Caption: message.Caption,
			})
		} else if message.Document != nil {
			fileID := message.Document.FileID
			api.SendDocument(echotron.NewInputFileID(fileID), feedbackChannelId, &echotron.DocumentOptions{
				Caption: message.Caption,
			})
		} else if message.Text != "" {
			api.SendMessage(message.Text, feedbackChannelId, nil)
		}
	}
}
