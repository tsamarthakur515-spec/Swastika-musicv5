# 🎵 Swastika Music V5 (Go + gogram)

High-quality Telegram music bot rewritten in **Go** using [gogram](https://github.com/amarnathcjd/gogram).

Port of the Python project: [SWASTIKA-MUSIC-V5](https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5).

---

## ✅ Current (Core)

| Command | Description |
|---------|-------------|
| `/start` | Welcome message |
| `/ping` | Latency check |
| `/help` | Command list |
| `/play <query\|url>` | Add song to queue / play |
| `/vplay <query\|url>` | Video play (queue) |
| `/pause` | Pause playback |
| `/resume` | Resume playback |
| `/skip` | Skip current track |
| `/stop` | Stop & clear queue |
| `/queue` | Show current queue |

> **Note:** Full Telegram Voice Chat streaming (gortc / ntgcalls) is next.  
> Right now the bot accepts commands, manages per-chat queue, and is ready for the stream layer.

---

## 🚀 Setup

1. Copy config:
   ```bash
   cp Config.env.example Config.env
   ```
2. Fill `API_ID`, `API_HASH`, `BOT_TOKEN`.
3. Run:
   ```bash
   go mod tidy
   go run .
   ```

### Docker / Railway
```bash
docker build -t swastika-music .
docker run --env-file Config.env swastika-music
```

On Railway set env vars: `API_ID`, `API_HASH`, `BOT_TOKEN`.

---

## 📁 Structure

```
├── main.go
├── internal/
│   ├── config/
│   ├── bot/
│   └── queue/
├── plugins/
│   ├── start.go
│   ├── play.go
│   └── controls.go
└── Dockerfile
```

---

## Roadmap

- [x] Core bot + queue + basic commands
- [ ] YouTube search / download (yt-dlp)
- [ ] Voice chat streaming (gortc or ntgcalls)
- [ ] Moderation / locks / games (from Python version)

MIT
