package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	APIID     int
	APIHash   string
	BotToken  string
	OwnerID   int64
	LogGroup  int64
}

func Load() (*Config, error) {
	apiIDStr := os.Getenv("API_ID")
	apiHash := os.Getenv("API_HASH")
	botToken := os.Getenv("BOT_TOKEN")

	if apiIDStr == "" || apiHash == "" || botToken == "" {
		return nil, fmt.Errorf("missing required env vars: API_ID, API_HASH, BOT_TOKEN")
	}

	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid API_ID: %w", err)
	}

	cfg := &Config{
		APIID:    apiID,
		APIHash:  apiHash,
		BotToken: botToken,
	}

	if v := os.Getenv("OWNER_ID"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.OwnerID = id
		}
	}
	if v := os.Getenv("LOG_GROUP_ID"); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.LogGroup = id
		}
	}

	return cfg, nil
}
