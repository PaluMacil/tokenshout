package api

// Player contains the data for a player
type Player struct {
	ID       int
	Position Coord
	Actions  []string
}
