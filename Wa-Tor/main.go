package main

func main() {
	gridSize := 20
	numFish := 50
	numSharks := 20
	grid := NewGrid(gridSize)
	grid.Initialize(numFish, numSharks)
	grid.Simulate(100)
}
