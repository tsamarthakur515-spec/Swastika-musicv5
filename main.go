package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env if present
	_ = godotenv.Load()

	apiID := os.Getenv("API_ID")
	apiHash := os.Getenv("API_HASH")
	botToken := os.Getenv("BOT_TOKEN")

	if apiID == "" || apiHash == "" || botToken == "" {
		log.Fatal("Missing required env: API_ID, API_HASH, BOT_TOKEN")
	}

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   mustAtoi(apiID),
		AppHash: apiHash,
	})
	if err != nil {
		log.Fatal("Failed to create client: ", err)
	}

	if err := client.Connect(); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	// Login as bot
	if _, err := client.LoginBot(botToken); err != nil {
		log.Fatal("Failed to login bot: ", err)
	}

	me, err := client.GetMe()
	if err != nil {
		log.Fatal("Failed to get me: ", err)
	}
	log.Printf("✅ Swastika Music V5 (Go/gogram) started as @%s", me.Username)

	// Basic /start handler
	client.AddMessageHandler(telegram.OnNewMessage, func(m *telegram.NewMessage) error {
		if m.IsCommand() && m.Command() == "start" {
			_, err := m.Reply("🎵 <b>Swastika Music V5</b>\n\nGo + gogram version.\nFull music features coming soon.\n\nUse /play, /vplay etc. (WIP)", telegram.SendOptions{
				ParseMode: telegram.HTML,
			})
			return err
		}
		return nil
	})

	// Keep running
	log.Println("Bot is running. Press Ctrl+C to stop.")
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	log.Println("Shutting down...")
}

func mustAtoi(s string) int {
	var n int
	for _, c := range s {
		if c < '0' || c > '9' {
			log.Fatal("Invalid API_ID")
		}
		n = n*10 + int(c-'0')
	}
	return n
}
