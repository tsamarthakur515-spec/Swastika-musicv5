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

	OwnerUsername   string
	SupportChat     string
	SupportChannel  string
	StartImageURL   string
	StatsImageURL   string
	PingImageURL    string

	YTAPIURL string
	YTAPIKey string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	TablePrefix string

	GroqAPIKey  string
	GroqAPIBase string
	GroqModel   string
}

func Load() (*Config, error) {
	apiIDStr := strings.TrimSpace(os.Getenv("API_ID"))
	apiHash := strings.TrimSpace(os.Getenv("API_HASH"))
	botToken := strings.TrimSpace(os.Getenv("BOT_TOKEN"))

	if apiIDStr == "" || apiHash == "" || botToken == "" {
		return nil, fmt.Errorf("missing required env vars: API_ID, API_HASH, BOT_TOKEN")
	}

	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid API_ID: %w", err)
	}

	cfg := &Config{
		APIID:           apiID,
		APIHash:         apiHash,
		BotToken:        botToken,
		StringSession:   strings.TrimSpace(os.Getenv("STRING_SESSION")),
		OwnerUsername:   strings.TrimPrefix(strings.TrimSpace(os.Getenv("OWNER_USERNAME")), "@"),
		SupportChat:     strings.TrimPrefix(strings.TrimSpace(os.Getenv("SUPPORT_CHAT")), "@"),
		SupportChannel:  strings.TrimPrefix(strings.TrimSpace(os.Getenv("SUPPORT_CHANNEL")), "@"),
		StartImageURL:   envOr("START_IMAGE_URL", "https://files.catbox.moe/mmr2ll.jpg"),
		StatsImageURL:   envOr("STATS_IMAGE_URL", "https://files.catbox.moe/hno0m2.jpg"),
		PingImageURL:    envOr("PING_IMAGE_URL", "https://files.catbox.moe/wfqfeh.jpg"),
		YTAPIURL:        envOr("API_URL", "https://aruyt.up.railway.app"),
		YTAPIKey:        envOr("API_KEY", "YUKI-XJd3KfUSWeuOWsiZyuIrlmQf"),
		DBHost:          strings.TrimSpace(os.Getenv("DB_HOST")),
		DBUser:          strings.TrimSpace(os.Getenv("DB_USER")),
		DBPassword:      strings.TrimSpace(os.Getenv("DB_PASSWORD")),
		DBName:          envOr("DB_NAME", "postgres"),
		TablePrefix:     envOr("TABLE_PREFIX", "swastika_"),
		GroqAPIKey:      strings.TrimSpace(os.Getenv("GROQ_API_KEY")),
		GroqAPIBase:     envOr("GROQ_API_BASE", "https://api.groq.com/openai/v1"),
		GroqModel:       envOr("GROQ_MODEL", "openai/gpt-oss-20b"),
	}

	if v := strings.TrimSpace(os.Getenv("OWNER_ID")); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.OwnerID = id
		}
	}
	if v := strings.TrimSpace(os.Getenv("LOG_GROUP_ID")); v != "" {
		if id, err := strconv.ParseInt(v, 10, 64); err == nil {
			cfg.LogGroup = id
		}
	}
	if v := strings.TrimSpace(os.Getenv("DB_PORT")); v != "" {
		if p, err := strconv.Atoi(v); err == nil {
			cfg.DBPort = p
		}
	} else {
		cfg.DBPort = 5432
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

func envOr(key, def string) string {
	if v := strings.TrimSpace(os.Getenv(key)); v != "" {
		return v
	}
	return def
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
