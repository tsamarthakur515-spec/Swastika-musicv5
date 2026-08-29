package database

// Database layer stub — port from Python PANDAMUSIC/modules/database.py
// Planned: PostgreSQL / SQLite for users, chats, sudo, welcomes, locks, stats.

type Store struct {
	// db *sql.DB
}

func New() *Store {
	return &Store{}
}

// TODO: users, chats, sudo list, welcome messages, lock settings, play history
