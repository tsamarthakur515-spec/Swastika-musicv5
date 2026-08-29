package main

import (
	"log"
	"os"
	"strconv"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env / Config.env if present
	_ = godotenv.Load()
	_ = godotenv.Load("Config.env")

	apiIDStr := os.Getenv("API_ID")
	apiHash := os.Getenv("API_HASH")
	botToken := os.Getenv("BOT_TOKEN")

	if apiIDStr == "" || apiHash == "" || botToken == "" {
		log.Fatal("Missing required env vars: API_ID, API_HASH, BOT_TOKEN. Set them in Config.env or environment.")
	}

	apiID, err := strconv.Atoi(apiIDStr)
	if err != nil {
		log.Fatal("Invalid API_ID: ", err)
	}

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   apiID,
		AppHash: apiHash,
		Session: "swastika.session",
	})
	if err != nil {
		log.Fatal("Failed to create client: ", err)
	}

	if err := client.Conn(); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	if err := client.LoginBot(botToken); err != nil {
		log.Fatal("Failed to login as bot: ", err)
	}

	me, err := client.GetMe()
	if err != nil {
		log.Fatal("Failed to get bot info: ", err)
	}
	log.Printf("✅ Swastika Music V5 (Go + gogram) started as @%s (ID: %d)\n", me.Username, me.ID)

	// /start command
	client.On("message:/start", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🎵 <b>Swastika Music V5</b>\n\n<code>Go + gogram</code> edition.\n\nFull music streaming, moderation, games & more are being ported from the Python version.\n\nStay tuned!", &telegram.SendOptions{
			ParseMode: "HTML",
		})
		return err
	})

	// Basic ping
	client.On("message:/ping", func(m *telegram.NewMessage) error {
		_, err := m.Reply("🏓 Pong! Bot is alive (Go).")
		return err
	})

	log.Println("Bot is running. Waiting for updates...")
	client.Idle()
}
