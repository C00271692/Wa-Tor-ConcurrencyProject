package main

type CellType int

const (
	Empty CellType = iota
	Fish
	Shark
)

type Cell struct {
	Type       CellType
	BreedTime  int
	StarveTime int
}
