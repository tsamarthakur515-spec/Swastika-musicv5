# 🎵 Swastika Music V5 (Go + gogram)

Full-surface Telegram music bot in **Go** using [gogram](https://github.com/amarnathcjd/gogram).

Port of [SWASTIKA-MUSIC-V5](https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5) (Python).

---

## ✅ Working now

| Area | Commands |
|------|----------|
| **Core** | `/start` `/help` `/ping` |
| **Music** | `/play` `/vplay` `/pause` `/resume` `/skip` `/stop` `/end` `/queue` `/replay` `/stream` |
| **Moderation** | `/ban` `/unban` `/mute` `/unmute` `/kick` `/promote` `/demote` |
| **Locks** | `/locks` `/lock` `/unlock` |
| **Welcome** | `/welcome` `/setwelcome` `/resetwelcome` |
| **Tag** | `/tagall` `/all` |
| **Fun** | `/chatbot` `/games` `/dice` `/truth` `/dare` `/guess` `/couple` `/crypto` `/price` |
| **Admin** | `/sudo` `/addsudo` `/rmsudo` `/sudolist` `/stats` `/gcast` `/broadcast` `/gban` `/ungban` |
| **Clone** | `/clone` `/clones` |
| **System** | `/maintenance` `/reload` `/restart` `/logger` `/active` `/activevc` `/track` `/noabuse` |

> **Music queue + controls work.**  
> Voice Chat streaming (gortc/ntgcalls), YouTube resolve, and full admin actions are the next implementation layer.

---

## Structure

```
├── main.go
├── internal/
│   ├── config/
│   ├── bot/
│   ├── queue/
│   ├── database/      (stub)
│   ├── helpers/
│   └── formatters/
├── platforms/
│   └── youtube.go     (stub)
├── plugins/           (all Python plugins mirrored)
│   ├── start.go  play.go  controls.go  stream.go
│   ├── moderation.go  locks.go  noabuse.go  welcome.go  tagall.go
│   ├── chatbot.go  games.go  guessgame.go  couple.go  crypto.go
│   ├── sudo.go  stats.go  gcast.go  clone.go  logger.go
│   ├── maintenance.go  reload.go  global.go  callbacks.go
│   ├── active.go  tracker.go
│   └── register.go
├── Dockerfile
└── Config.env.example
```

---

## Setup

```bash
cp Config.env.example Config.env
# fill API_ID, API_HASH, BOT_TOKEN
go mod tidy
go run .
```

**Railway:** set `API_ID`, `API_HASH`, `BOT_TOKEN` → deploy.

---

## Roadmap

1. ✅ Core bot + full plugin surface
2. ⏳ YouTube search (yt-dlp)
3. ⏳ VC streaming (gortc / ntgcalls)
4. ⏳ Database (sudo, welcome, locks, stats)
5. ⏳ Full moderation / games logic from Python

MIT
