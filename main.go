package main

import (
	"pu/bot"
	"pu/internal/config"
)

func main() {
	cfg := config.MustLoad()

	bot.Run(cfg.Token, cfg.Username, cfg.ChatID)
}
