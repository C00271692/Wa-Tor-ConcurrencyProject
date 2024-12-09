package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"image/color"
	"time"
)

func RunGUI(grid *Grid) {
	a := app.New()
	w := a.NewWindow("Wa-Tor Simulation")
	w.Resize(fyne.NewSize(400, 400))

	canvas := canvas.NewRasterWithPixels(grid.Draw)
	w.SetContent(container.NewWithoutLayout(canvas))

	go func() {
		for i := 0; i < 100; i++ {
			grid.Update()
			canvas.Refresh()
			time.Sleep(1 * time.Second)
		}
	}()

	w.ShowAndRun()
}

func (g *Grid) Draw(x, y, w, h int) color.Color {
	if w == 0 || h == 0 {
		return color.RGBA{255, 255, 255, 255}
	}

	cellSize := w / g.Size
	gridX, gridY := x/cellSize, y/cellSize

	if gridX >= g.Size || gridY >= g.Size {
		return color.RGBA{255, 255, 255, 255}
	}

	switch g.Cells[gridX][gridY].Type {
	case Empty:
		return color.RGBA{255, 255, 255, 255}
	case Fish:
		return color.RGBA{0, 255, 0, 255}
	case Shark:
		return color.RGBA{255, 0, 0, 255}
	}
	return color.RGBA{255, 255, 255, 255}
}
