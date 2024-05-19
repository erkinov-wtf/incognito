package command

import (
	"github.com/NicoNex/echotron/v3"
	"pu/pkg/replies/commands"
	"strings"
	"sync"
)

// HandleCommand processes bot commands
func HandleCommand(api *echotron.API, message *echotron.Message, feedbackChannelId int64, feedbackUsers map[int64]bool, mu *sync.Mutex) {
	// Extract the command from the message text
	command := strings.Split(message.Text, " ")[0]

	mu.Lock()
	defer mu.Unlock()

	switch command {
	case "/start":
		api.SendMessage(commands.StartMessage, message.Chat.ID, nil)
	case "/feedback":
		feedbackUsers[message.Chat.ID] = true
		api.SendMessage(commands.FeedbackMessage, message.Chat.ID, nil)
	case "/issues":
		api.SendMessage(commands.KnownIssues, message.Chat.ID, nil)
	case "/help":
		api.SendMessage(commands.HelpMessage, message.Chat.ID, nil)
	default:
		api.SendMessage(commands.UnknownCommandMessage, message.Chat.ID, nil)
	}
}
