package api

// Coord is an x, y coordinate
type Coord struct {
	X int
	Y int
}

// Map indicated map size and rules
type Map struct {
	Size  Size
	Rules []string
	Grid  Grid `json:"-"`
}

// Grid holds the two dimensional slice of rooms
type Grid [][]Room

// Size indicates width and height
type Size struct {
	Width  int
	Height int
}
