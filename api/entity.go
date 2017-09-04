package api

// Entity is the struct containing a player, token, item (future) or monster (future)
type Entity struct {
	Type string
	ID   int
}
