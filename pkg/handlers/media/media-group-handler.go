package media

import (
	"github.com/NicoNex/echotron/v3"
	"pu/pkg/types"
	"sync"
	"time"
)

func HandleMediaGroup(api *echotron.API, message *echotron.Message, chatID int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, mu *sync.Mutex, mediaType string) {
	var media echotron.GroupableInputMedia

	if mediaType == types.Photo {
		media = echotron.InputMediaPhoto{
			Type:    types.Photo,
			Media:   echotron.NewInputFileID(message.Photo[len(message.Photo)-1].FileID),
			Caption: message.Caption,
		}
	} else if mediaType == types.Doc {
		media = echotron.InputMediaDocument{
			Type:    types.Doc,
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
		if mediaType == types.Photo {
			photoID := message.Photo[len(message.Photo)-1].FileID
			api.SendPhoto(echotron.NewInputFileID(photoID), chatID, &echotron.PhotoOptions{
				Caption: message.Caption,
			})
		} else if mediaType == types.Doc {
			fileID := message.Document.FileID
			api.SendDocument(echotron.NewInputFileID(fileID), chatID, &echotron.DocumentOptions{
				Caption: message.Caption,
			})
		}
	}
}
