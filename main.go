package main

import (
	"github.com/NicoNex/echotron/v3"
	"log"
	"pu/handlers/message"
	"pu/internal/config"
	"sync"
)

func main() {
	cfg := config.MustLoad()
	api := echotron.NewAPI(cfg.Token)
	feedbackUsers := make(map[int64]bool)

	log.Println("App started")

	mediaGroupMap := make(map[string][]echotron.GroupableInputMedia)
	originalMessages := make(map[int64]int)
	var mu sync.Mutex

	// Listen for messages
	for u := range echotron.PollingUpdates(cfg.Token) {
		log.Printf("Received update: %+v\n", u)
		if u.Message != nil {
			message.HandleMessage(&api, u.Message, cfg.ChannelID, &mediaGroupMap, &originalMessages, &mu, cfg.FeedbackChannelId, feedbackUsers)
		} else {
			log.Println("Received update with no message field")
		}
	}
}
