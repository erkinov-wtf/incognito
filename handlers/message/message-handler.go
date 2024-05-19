package message

import (
	"fmt"
	"pu/pkg/handlers"
	"pu/pkg/handlers/command"
	"pu/pkg/handlers/feedback"
	"sync"

	"github.com/NicoNex/echotron/v3"
	"pu/pkg/handlers/media"
	"pu/pkg/types"
)

func HandleMessage(api *echotron.API, message *echotron.Message, chatID int64, mediaGroupMap *map[string][]echotron.GroupableInputMedia, originalMessages *map[int64]int, mu *sync.Mutex, feedbackChannelId int64, feedbackUsers map[int64]bool) {
	if message == nil {
		fmt.Println("Received nil message")
		return
	}

	mu.Lock()
	if inFeedbackMode, ok := feedbackUsers[message.Chat.ID]; ok && inFeedbackMode {
		// Remove the user from feedback mode after receiving feedback
		delete(feedbackUsers, message.Chat.ID)
		mu.Unlock()
		feedback.SendFeedback(api, message, feedbackChannelId, mediaGroupMap, mu)
		api.SendMessage("Your feedback has been sent. Thank you!", message.Chat.ID, nil)
		return
	}
	mu.Unlock()

	// Check if the message is a command
	if message.Text != "" && message.Text[0] == '/' {
		command.HandleCommand(api, message, feedbackChannelId, feedbackUsers, mu)
		return
	}

	switch {
	case message.Text != "":
		// If the message is a text, send it to the channel without sender information
		sentMsg, err := api.SendMessage(message.Text, chatID, nil)
		if err == nil && sentMsg.Result != nil {
			mu.Lock()
			(*originalMessages)[int64(message.ID)] = sentMsg.Result.ID
			mu.Unlock()
		} else {
			fmt.Println("Error sending message or sentMsg.Result is nil:", err)
		}
	case message.Photo != nil:
		// If the message is a photo, handle media group
		media.HandleMediaGroup(api, message, chatID, mediaGroupMap, mu, types.Photo)
	case message.Document != nil:
		// If the message is a document, handle media group
		media.HandleMediaGroup(api, message, chatID, mediaGroupMap, mu, types.Doc)
	}

	// Check if message.ExternalReply is nil before accessing it
	//if message.ExternalReply != nil {
	//	fmt.Println(message.ExternalReply)
	//	replies.HandleReplyMessage(api, message, chatID, originalMessages, mu)
	//}

	// Reply to the user
	handlers.ReplyMessage(api, message)
}
