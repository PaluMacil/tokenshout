package api

// Room defines a single room in the game
type Room struct {
	Doors          []string
	Entities       []Entity
	AllowedActions []string
}
