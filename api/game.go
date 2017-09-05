package api

import uuid "github.com/satori/go.uuid"

// UpdateTick contains changes in gamestate contained in an update tick
type UpdateTick struct {
	GameID         uuid.UUID
	Players        []Player
	Room           Room
	AllowedActions []string
}

// Game contains the gamestate
type Game struct {
	GameID  uuid.UUID
	Map     Map
	Players []Player `json:"-"`
	Version int
}

// WelcomeMessage contains all the data a player needs at the start
// of a game.
type WelcomeMessage struct {
	Game
	PlayerCount int
	YourID      int
}
