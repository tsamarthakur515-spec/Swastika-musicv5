package botapi

// Extra Bot API helpers — port from Python modules/bot_api.py
// Thin wrappers around common admin / chat actions via gogram.

import (
	"github.com/amarnathcjd/gogram/telegram"
)

type API struct {
	Client *telegram.Client
}

func New(c *telegram.Client) *API {
	return &API{Client: c}
}

// TODO: ban, unban, mute, promote, get_chat_member, restrict wrappers
