package main

import (
	"log"

	"github.com/amarnathcjd/gogram/telegram"
	"github.com/joho/godotenv"

	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/config"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/console"
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/stream"
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

	// VC streaming: prefer userbot session for gortc join
	var streamClient *telegram.Client
	vcEnabled := false
	if cfg.StringSession != "" {
		uc, err := telegram.NewClient(telegram.ClientConfig{
			AppID:         cfg.APIID,
			AppHash:       cfg.APIHash,
			StringSession: cfg.StringSession,
			Session:       "assistant.session",
		})
		if err == nil && uc.Conn() == nil {
			streamClient = uc
			vcEnabled = true
			log.Println("Assistant userbot connected — VC stream enabled")
		} else {
			log.Printf("Assistant login failed: %v — VC disabled, file mode only", err)
		}
	} else {
		log.Println("No STRING_SESSION — VC stream disabled (file delivery only)")
	}

	streamMgr := stream.New(streamClient, vcEnabled)
	b := bot.New(client, cfg, streamMgr)
	plugins.Register(b)

	log.Println("Bot is running. Waiting for updates...")
	client.Idle()
}
