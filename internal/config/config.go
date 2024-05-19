package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
)

type Config struct {
	Token             string `env:"TOKEN"`
	Username          string `env:"USERNAME"`
	ChannelID         int64  `env:"CHANNEL_ID"`
	FeedbackChannelId int64  `env:"FEEDBACK_CHANNEL_ID"`
}

func MustLoad() *Config {
	configPath := ".env"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config file: %s", err)
	}

	return &cfg
}
