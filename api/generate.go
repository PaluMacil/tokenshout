package api

// ProofOfConceptMap generates a proof of concept map which is
// always the same static map
func ProofOfConceptMap(playerCount int) Map {
	grid := Grid{
		{Room{}, Room{}, Room{}, Room{}, Room{}},
		{Room{}, Room{}, Room{}, Room{}, Room{}},
		{Room{}, Room{}, Room{}, Room{}, Room{}},
		{Room{}, Room{}, Room{}, Room{}, Room{}},
		{Room{}, Room{}, Room{}, Room{}, Room{}},
	}
	return Map{
		Grid:  grid,
		Rules: []string{"Doors", "Checkpoints", "Traps", "WinCheckpoints"},
		Size:  Size{Width: 5, Height: 5},
	}
}
