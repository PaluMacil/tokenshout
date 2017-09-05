package api

// Coord is an x, y coordinate
type Coord struct {
	X int
	Y int
}

// Grid holds the two dimensional slice of rooms
type Grid [][]Room

// Map indicated map size and rules
type Map struct {
	Size  Size
	Rules []string
	Grid  Grid `json:"-"`
}

// Room defines a single room in the game
type Room struct {
	Doors          []string
	Entities       []Entity
	AllowedActions []string
}

// Size indicates width and height
type Size struct {
	Width  int
	Height int
}
