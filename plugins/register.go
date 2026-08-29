package plugins

import (
	"github.com/tsamarthakur515-spec/Swastika-musicv5/internal/bot"
)

// Register wires every plugin (core + full Python-port surface).
func Register(b *bot.Bot) {
	// Core music
	registerStart(b)
	registerPlay(b)
	registerControls(b)
	registerStream(b)

	// Group management
	registerModeration(b)
	registerLocks(b)
	registerNoAbuse(b)
	registerWelcome(b)
	registerTagAll(b)

	// Fun / social
	registerChatbot(b)
	registerGames(b)
	registerGuessGame(b)
	registerCouple(b)
	registerCrypto(b)
	registerRob(b)

	// Admin / system
	registerSudo(b)
	registerStats(b)
	registerGcast(b)
	registerClone(b)
	registerLogger(b)
	registerMaintenance(b)
	registerReload(b)
	registerGlobal(b)
	registerCallbacks(b)
	registerActive(b)
	registerTracker(b)
	registerKill(b)
}
