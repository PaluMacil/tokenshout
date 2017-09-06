package api

var allDoors = []string{"N", "E", "S", "W"}
var neDoors = []string{"N", "E"}
var nwDoors = []string{"N", "W"}
var seDoors = []string{"E", "S"}
var swDoors = []string{"S", "W"}
var nesDoors = []string{"N", "E", "S"}
var nwsDoors = []string{"N", "W", "S"}
var eswDoors = []string{"E", "S", "W"}
var enwDoors = []string{"E", "n", "W"}

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
					grid[x][y].Doors = seDoors
				} else if y == 4 {
					grid[x][y].Doors = neDoors
				} else {
					grid[x][y].Doors = nesDoors
				}
			} else if x == 4 {
				//generate south edge and south corners
				if y == 0 {
					grid[x][y].Doors = swDoors
				} else if y == 4 {
					grid[x][y].Doors = nwDoors
				} else {
					grid[x][y].Doors = nwsDoors
				}
			} else if y == 0 {
				//generate north edge
				grid[x][y].Doors = eswDoors
			} else if y == 4 {
				//generate south edge
				grid[x][y].Doors = enwDoors
			} else {
				grid[x][y].Doors = allDoors
			}
		}
		start := Entity{Type: "Start", ID: 0}
		tokenOne := Entity{Type: "Token", ID: 1}
		tokenTwo := Entity{Type: "Token", ID: 2}
		end := Entity{Type: "End", ID: 3}
		grid[1][4].Entities = append(grid[1][4].Entities, start)
		grid[3][4].Entities = append(grid[3][4].Entities, tokenOne)
		grid[0][1].Entities = append(grid[0][1].Entities, tokenTwo)
		grid[4][2].Entities = append(grid[4][2].Entities, end)
		for p := 0; p < playerCount; p++ {
			player := Entity{
				Type: "Player",
				ID:   p,
			}
			grid[1][4].Entities = append(grid[1][4].Entities, player)
		}
	}
	return Map{
		Grid:  grid,
		Rules: []string{"Doors", "Checkpoints", "Traps", "WinCheckpoints"},
		Size:  Size{Width: 5, Height: 5},
	}
}
