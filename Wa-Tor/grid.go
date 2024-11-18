package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Grid struct {
	Cells [][]Cell
	Size  int
}

// define a Grid structure
func NewGrid(size int) *Grid {
	grid := &Grid{
		Cells: make([][]Cell, size),
		Size:  size,
	}
	for i := range grid.Cells {
		grid.Cells[i] = make([]Cell, size)
	}
	return grid
}

//NewGrid creates a new Grid of the given size and initializes its cells.

func (g *Grid) Initialize(numFish, numSharks int) {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator using current time.

	// Place fish randomly on the grid.
	for i := 0; i < numFish; i++ {
		x, y := rand.Intn(g.Size), rand.Intn(g.Size)
		g.Cells[x][y] = Cell{Type: Fish}
	}

	// Place sharks randomly on the grid.
	for i := 0; i < numSharks; i++ {
		x, y := rand.Intn(g.Size), rand.Intn(g.Size)
		g.Cells[x][y] = Cell{Type: Shark, StarveTime: 5} // Example starve time for sharks, can tweak or change values. see what happens?.
	}
}

func (g *Grid) Update() {
	// Iterate over each cell in the grid
	for x := 0; x < g.Size; x++ {
		for y := 0; y < g.Size; y++ {
			// Check the type of the cell and move the fish or shark accordingly
			switch g.Cells[x][y].Type {
			case Fish:
				g.MoveFish(x, y) // Move the fish
			case Shark:
				g.MoveShark(x, y)
			}
		}
	}
}

func (g *Grid) Simulate(steps int) {
	for i := 0; i < steps; i++ {
		g.Update()
		g.Display()
		time.Sleep(1 * time.Second) // Pause for a second between steps
	}
}

func (g *Grid) Display() {
	//row
	for x := 0; x < g.Size; x++ {
		//column
		for y := 0; y < g.Size; y++ {
			switch g.Cells[x][y].Type {
			//change to colours when implement GUI
			case Empty:
				fmt.Print(".")
			case Fish:
				fmt.Print("F")
			case Shark:
				fmt.Print("S")
			}
		}
		//newline between steps
		fmt.Println()
	}
	fmt.Println()
}
