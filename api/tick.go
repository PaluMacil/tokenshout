package api

import uuid "github.com/satori/go.uuid"

// UpdateTick contains changes in gamestate contained in an update tick
type UpdateTick struct {
	GameID  uuid.UUID
	Players []Player
}
