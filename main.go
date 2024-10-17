package main

import (
	"fmt"

	// for gui

	"github.com/veandco/go-sdl2/sdl"
)

const (
	winWidth  = 800
	winHeight = 600
)

var cube [5][5][5]int = [5][5][5]int{
	{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}},
	{{26, 27, 28, 29, 30}, {31, 32, 33, 34, 35}, {36, 37, 38, 39, 40}, {41, 42, 43, 44, 45}, {46, 47, 48, 49, 50}},
	{{51, 52, 53, 54, 55}, {56, 57, 58, 59, 60}, {61, 62, 63, 64, 65}, {66, 67, 68, 69, 70}, {71, 72, 73, 74, 75}},
	{{76, 77, 78, 79, 80}, {81, 82, 83, 84, 85}, {86, 87, 88, 89, 90}, {91, 92, 93, 94, 95}, {96, 97, 98, 99, 100}},
	{{101, 102, 103, 104, 105}, {106, 107, 108, 109, 110}, {111, 112, 113, 114, 115}, {116, 117, 118, 119, 120}, {121, 122, 123, 124, 125}},
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("Error initializing SDL: %v\n", err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Go OpenGL Buttons", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Error creating window: %v\n", err)
		return
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		fmt.Printf("Error getting window surface: %v\n", err)
		return
	}

	buttons := []sdl.Rect{
		{X: 50, Y: 50, W: 200, H: 100},   // Button 1
		{X: 300, Y: 50, W: 200, H: 100},  // Button 2
		{X: 50, Y: 200, W: 200, H: 100},  // Button 3
		{X: 300, Y: 200, W: 200, H: 100}, // Button 4
	}

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseButtonEvent:
				if t.Type == sdl.MOUSEBUTTONDOWN {
					for i, button := range buttons {
						if t.X >= button.X && t.X <= button.X+button.W && t.Y >= button.Y && t.Y <= button.Y+button.H {
							switch i {
							case 0:
								fmt.Println("Button 1 pressed: Executing search.Singular() and displaying cube")
								drawCube(surface)
							case 1:
								fmt.Println("Button 2 pressed: Executing search.Multiple()")
								// search.Multiple()
							case 2:
								fmt.Println("Button 3 pressed: Executing search.AllAlgorithms()")
								// search.AllAlgorithms()
							case 3:
								fmt.Println("Button 4 pressed: Executing search.CheckOF()")
								// search.CheckOF()
							}
						}
					}
				}
			}
		}

		surface.FillRect(nil, 0x000000) // Fill the screen with black
		for _, button := range buttons {
			surface.FillRect(&button, 0x00FF00) // Draw buttons in green
		}
		window.UpdateSurface()
	}
}

func drawCube(surface *sdl.Surface) {
	// Clear the surface
	surface.FillRect(nil, 0x000000)

	// Draw the cube representation
	cellSize := 20
	startX := 50
	startY := 350
	for z := 0; z < 5; z++ {
		for y := 0; y < 5; y++ {
			for x := 0; x < 5; x++ {
				rect := sdl.Rect{
					X: int32(startX + x*cellSize + z*cellSize),
					Y: int32(startY + y*cellSize + z*cellSize),
					W: int32(cellSize - 2),
					H: int32(cellSize - 2),
				}
				surface.FillRect(&rect, 0x00FFFF) // Draw the cube cell in cyan
			}
		}
	}
}
