package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	APIID         int
	APIHash       string
	BotToken      string
	OwnerID       int64
	LogGroup      int64
	StringSession string
	SudoUsers     []int64
	// YouTube helpers (Python Youtube.py parity)
	YTAPIURL string
	YTAPIKey string
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
		APIID:         apiID,
		APIHash:       apiHash,
		BotToken:      botToken,
		StringSession: os.Getenv("STRING_SESSION"),
		YTAPIURL:      os.Getenv("API_URL"),
		YTAPIKey:      os.Getenv("API_KEY"),
	}

	// Defaults from Python PANDAMUSIC/platforms/Youtube.py
	if cfg.YTAPIURL == "" {
		cfg.YTAPIURL = "https://aruyt.up.railway.app"
	}
	if cfg.YTAPIKey == "" {
		cfg.YTAPIKey = "YUKI-XJd3KfUSWeuOWsiZyuIrlmQf"
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
	if v := os.Getenv("SUDO_USERS"); v != "" {
		for _, p := range strings.Split(v, ",") {
			p = strings.TrimSpace(p)
			if id, err := strconv.ParseInt(p, 10, 64); err == nil {
				cfg.SudoUsers = append(cfg.SudoUsers, id)
			}
		}
	}

	return cfg, nil
}

func (c *Config) IsSudo(userID int64) bool {
	if c.OwnerID != 0 && userID == c.OwnerID {
		return true
	}
	for _, id := range c.SudoUsers {
		if id == userID {
			return true
		}
	}
	return false
}
