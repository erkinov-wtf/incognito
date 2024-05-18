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

	log.Println("App started")

	mediaGroupMap := make(map[string][]echotron.GroupableInputMedia)
	var mu sync.Mutex

	// Listen for messages
	for u := range echotron.PollingUpdates(cfg.Token) {
		if u.Message != nil {
			message.HandleMessage(&api, u.Message, cfg.ChannelID, &mediaGroupMap, &mu)
		}
	}
}
