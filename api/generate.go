package api

var allDoors = []string{"N", "E", "S", "W"}
var neDoors = []string{"N", "E"}
var nwDoors = []string{"N", "W"}
var seDoors = []string{"E", "S"}
var swDoors = []string{"S", "W"}
var allActions = []string{
	"Move",
	"Shout",
	"Wait",
}

// ProofOfConceptMap generates a proof of concept map which is
// always the same static map
func ProofOfConceptMap(playerCount int) Map {
	var grid Grid
	/*
		0,0  1,0  2,0  3,0  4,0
		0,1  1,1  2,1  3,1  4,1
		0,2  1,2  2,2  3,2  4,2
		0,3  1,3  2,3  3,3  4,3
		0,4  1,4  2,4  3,4  4,4
	*/
	for x := 0; x < 5; x++ {
		//for each x, create a column of rooms
		col := make([]Room, 5, 5)
		grid = append(grid, col)
		for y := 0; y < 5; y++ {
			grid[x][y].AllowedActions = allActions
			if x == 0 {
				//generate north edge and north corners
				if y == 0 {

				} else if y == 4 {

				} else {

				}
			} else if x == 4 {
				//generate south edge and south corners
				if y == 0 {

				} else if y == 4 {

				} else {

				}
			} else if y == 0 {
				//generate west edge

			} else if y == 4 {
				//generate east edge

			} else {
				grid[x][y].Doors = allDoors
			}
		}
	}
	return Map{
		Grid:  grid,
		Rules: []string{"Doors", "Checkpoints", "Traps", "WinCheckpoints"},
		Size:  Size{Width: 5, Height: 5},
	}
}
