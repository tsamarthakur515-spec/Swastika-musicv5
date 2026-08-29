package main

import (
	"log"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/joho/godotenv"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/config"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/console"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/platforms"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/plugins"
)

func main() {
	_ = godotenv.Load()
	_ = godotenv.Load("Config.env")

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	platforms.Configure(cfg.YTAPIURL, cfg.YTAPIKey)

	client, err := telegram.NewClient(telegram.ClientConfig{
		AppID:   cfg.APIID,
		AppHash: cfg.APIHash,
		Session: "swastika.session",
	})
	if err != nil {
		log.Fatal("Failed to create client: ", err)
	}

	if err := client.Conn(); err != nil {
		log.Fatal("Failed to connect: ", err)
	}

	if err := client.LoginBot(cfg.BotToken); err != nil {
		log.Fatal("Failed to login as bot: ", err)
	}

	me, err := client.GetMe()
	if err != nil {
		log.Fatal("Failed to get bot info: ", err)
	}

	console.Banner(me.Username, me.ID)

	b := bot.New(client, cfg)
	plugins.Register(b)

	log.Println("Bot is running. Waiting for updates...")
	client.Idle()
}
