# 🎵 Swastika Music V5 (Go + gogram)

Go port of [SWASTIKA-MUSIC-V5](https://github.com/tsamarthakur515-spec/SWASTIKA-MUSIC-V5).

## YouTube layer (same as Python)

| Step | Python | Go (this repo) |
|------|--------|----------------|
| Search | youtubesearchpython → yt-dlp | yt-dlp `ytsearch` |
| Download | **API** (`API_URL`/`API_KEY`) first | **Same API** first |
| Fallback | local yt-dlp | local yt-dlp |
| Cache | `downloads/` | `downloads/` |

Defaults (same as Python `Youtube.py`):

```
API_URL=https://aruyt.up.railway.app
API_KEY=YUKI-XJd3KfUSWeuOWsiZyuIrlmQf
```

## Working

- `/play` `/vplay` — search + API/yt-dlp download + send media
- Queue, pause, skip, stop, etc.
- Full plugin command surface

## Not yet (Python has, Go next)

- **Voice Chat streaming** via ntgcalls / py-tgcalls  
  → needs gortc/ntgcalls + often a userbot `STRING_SESSION`

## Railway env

`API_ID` `API_HASH` `BOT_TOKEN`  
Optional: `API_URL` `API_KEY` `OWNER_ID` `LOG_GROUP_ID`

MIT
