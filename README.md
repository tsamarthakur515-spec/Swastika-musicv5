# 🎵 Swastika Music V5 (Go + gogram)

Telegram music bot in **Go** + [gogram](https://github.com/amarnathcjd/gogram).

Port of [SWASTIKA-MUSIC-V5](https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5).

---

## ✅ Working now

- `/play` `/vplay` — **YouTube search + resolve via yt-dlp**, queue, download MP3, send audio to chat
- `/pause` `/resume` `/skip` `/stop` `/queue` `/replay`
- Full plugin command surface (moderation, games, sudo, …)

> **Voice Chat (group call) streaming** is next (gortc / ntgcalls).  
> Right now `/play` sends the track as an **audio file** in the chat.

---

## Setup

```bash
cp Config.env.example Config.env
# API_ID, API_HASH, BOT_TOKEN
go mod tidy && go run .
```

**Needs:** `yt-dlp` + `ffmpeg` on PATH (Dockerfile installs both).

**Railway env:** `API_ID`, `API_HASH`, `BOT_TOKEN`

---

## Docker

```bash
docker build -t swastika-music .
docker run --env-file Config.env swastika-music
```

---

## Roadmap

1. ✅ Plugin surface + queue
2. ✅ YouTube via yt-dlp + audio file delivery
3. ⏳ Telegram Voice Chat streaming (gortc)
4. ⏳ Database (sudo, welcome, locks)
5. ⏳ Full moderation / games logic

MIT
