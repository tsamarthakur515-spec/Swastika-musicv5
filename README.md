# 🎵 Swastika Music V5 (Go + gogram)

**High quality Telegram music bot** — rewritten in **Go** using [gogram](https://github.com/amarnathcjd/gogram).

This is the Go port of the original Python project: [SWASTIKA-MUSIC-V5](https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5).

> ⚠️ **Status**: Core bot skeleton is ready. Full music streaming (ntgcalls), plugins (moderation, games, chatbot, etc.) are WIP and will be ported progressively.

---

## ✨ Current Features

- Bot login via `BOT_TOKEN`
- `/start` and `/ping` commands
- Config via `Config.env` or environment variables
- Docker-ready

## 🚀 Setup

1. Copy example config:
   ```bash
   cp Config.env.example Config.env
   ```
2. Fill `API_ID`, `API_HASH`, `BOT_TOKEN` in `Config.env` (get from my.telegram.org + @BotFather).
3. Run locally:
   ```bash
   go mod tidy
   go run .
   ```

Or with Docker:
```bash
docker build -t swastika-music .
docker run --env-file Config.env swastika-music
```

## 📦 Deployment notes

- This is a **Go** project (`go.mod` present).  
- If your hosting panel still shows **"No Python file found in repository!"**, change the **runtime / language** of the service from Python to **Go** (or Docker) in the dashboard.  
- Do **not** keep a Python buildpack/runtime selected.

## 🔗 Original Python version

All features (play/vplay, queue, moderation, locks, games, chatbot, clone bots, etc.) exist in:
https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5

They will be re-implemented here in idiomatic Go + gogram + ntgcalls.

## License

MIT — same spirit as the Python project.
