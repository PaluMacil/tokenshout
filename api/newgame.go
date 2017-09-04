package api

import uuid "github.com/satori/go.uuid"

// NewGame contains all the data a player needs at the start
// of a game.
type NewGame struct {
	GameID  uuid.UUID
	Map     Map
	Players int
	YourID  int
	Version int
}
