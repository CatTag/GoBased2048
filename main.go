package main

import (
	//"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand/v2"
	"strconv"
)

func emulateInputs(gameGrid [4][4]int) bool {
	initGameGrid := gameGrid
	for x := range len(gameGrid) {
		for y := range len(gameGrid[x]) - 1 {
			if gameGrid[x][y] == 0 {
				continue
			}
			if gameGrid[x][y+1] == 0 {
				gameGrid[x][y+1] = gameGrid[x][y]
				gameGrid[x][y] = 0
				for y2 := y; y2 > 0; y2-- {
					gameGrid[x][y2] = gameGrid[x][y2-1]
				}
				gameGrid[x][0] = 0
			} else if gameGrid[x][y+1] == gameGrid[x][y] {
				gameGrid[x][y+1] *= 2
				for y2 := y; y2 > 0; y2-- {
					gameGrid[x][y2] = gameGrid[x][y2-1]
				}
				gameGrid[x][0] = 0
			}
		}
	}
	if initGameGrid != gameGrid {
		return false
	}

	for x := range len(gameGrid) {
		for y := len(gameGrid[x]) - 1; y > 0; y-- {
			if gameGrid[x][y] == 0 {
				continue
			}
			if gameGrid[x][y-1] == 0 {
				gameGrid[x][y-1] = gameGrid[x][y]
				gameGrid[x][y] = 0
				for y2 := y; y2 < len(gameGrid[x])-1; y2++ {
					gameGrid[x][y2] = gameGrid[x][y2+1]
				}
				gameGrid[x][len(gameGrid[x])-1] = 0
			} else if gameGrid[x][y-1] == gameGrid[x][y] {
				gameGrid[x][y-1] *= 2
				for y2 := y; y2 < len(gameGrid[x])-1; y2++ {
					gameGrid[x][y2] = gameGrid[x][y2+1]
				}
				gameGrid[x][len(gameGrid[x])-1] = 0
			}
		}
	}
	if initGameGrid != gameGrid {
		return false
	}
	for y := range len(gameGrid) {
		for x := range len(gameGrid[y]) - 1 {
			if gameGrid[x][y] == 0 {
				continue
			}
			if gameGrid[x+1][y] == 0 {
				gameGrid[x+1][y] = gameGrid[x][y]
				gameGrid[x][y] = 0
				for x2 := x; x2 > 0; x2-- {
					gameGrid[x2][y] = gameGrid[x2-1][y]
				}
				gameGrid[0][y] = 0
			} else if gameGrid[x+1][y] == gameGrid[x][y] {
				gameGrid[x+1][y] *= 2
				for x2 := x; x2 > 0; x2-- {
					gameGrid[x2][y] = gameGrid[x2-1][y]
				}
				gameGrid[0][y] = 0
			}
		}
	}
	if initGameGrid != gameGrid {
		return false
	}
	for y := range len(gameGrid) {
		for x := len(gameGrid[y]) - 1; x > 0; x-- {
			if gameGrid[x][y] == 0 {
				continue
			}
			if gameGrid[x-1][y] == 0 {
				gameGrid[x-1][y] = gameGrid[x][y]
				gameGrid[x][y] = 0
				for x2 := x; x2 < len(gameGrid[y])-1; x2++ {
					gameGrid[x2][y] = gameGrid[x2+1][y]
				}
				gameGrid[len(gameGrid[y])-1][y] = 0
			} else if gameGrid[x-1][y] == gameGrid[x][y] {
				gameGrid[x-1][y] *= 2
				for x2 := x; x2 < len(gameGrid[y])-1; x2++ {
					gameGrid[x2][y] = gameGrid[x2+1][y]
				}
				gameGrid[len(gameGrid[y])-1][y] = 0
			}
		}
	}
	if initGameGrid != gameGrid {
		return false
	}
	return true

}

func runInputs(gameGrid [4][4]int) [4][4]int {
	if rl.IsKeyPressed(rl.KeyDown) {
		for x := range len(gameGrid) {
			for y := range len(gameGrid[x]) - 1 {
				if gameGrid[x][y] == 0 {
					continue
				}
				if gameGrid[x][y+1] == 0 {
					gameGrid[x][y+1] = gameGrid[x][y]
					gameGrid[x][y] = 0
					for y2 := y; y2 > 0; y2-- {
						gameGrid[x][y2] = gameGrid[x][y2-1]
					}
					gameGrid[x][0] = 0
				} else if gameGrid[x][y+1] == gameGrid[x][y] {
					gameGrid[x][y+1] *= 2
					for y2 := y; y2 > 0; y2-- {
						gameGrid[x][y2] = gameGrid[x][y2-1]
					}
					gameGrid[x][0] = 0
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyUp) {
		for x := range len(gameGrid) {
			for y := len(gameGrid[x]) - 1; y > 0; y-- {
				if gameGrid[x][y] == 0 {
					continue
				}
				if gameGrid[x][y-1] == 0 {
					gameGrid[x][y-1] = gameGrid[x][y]
					gameGrid[x][y] = 0
					for y2 := y; y2 < len(gameGrid[x])-1; y2++ {
						gameGrid[x][y2] = gameGrid[x][y2+1]
					}
					gameGrid[x][len(gameGrid[x])-1] = 0
				} else if gameGrid[x][y-1] == gameGrid[x][y] {
					gameGrid[x][y-1] *= 2
					for y2 := y; y2 < len(gameGrid[x])-1; y2++ {
						gameGrid[x][y2] = gameGrid[x][y2+1]
					}
					gameGrid[x][len(gameGrid[x])-1] = 0
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyRight) {
		for y := range len(gameGrid) {
			for x := range len(gameGrid[y]) - 1 {
				if gameGrid[x][y] == 0 {
					continue
				}
				if gameGrid[x+1][y] == 0 {
					gameGrid[x+1][y] = gameGrid[x][y]
					gameGrid[x][y] = 0
					for x2 := x; x2 > 0; x2-- {
						gameGrid[x2][y] = gameGrid[x2-1][y]
					}
					gameGrid[0][y] = 0
				} else if gameGrid[x+1][y] == gameGrid[x][y] {
					gameGrid[x+1][y] *= 2
					for x2 := x; x2 > 0; x2-- {
						gameGrid[x2][y] = gameGrid[x2-1][y]
					}
					gameGrid[0][y] = 0
				}
			}
		}
	}
	if rl.IsKeyPressed(rl.KeyLeft) {
		for y := range len(gameGrid) {
			for x := len(gameGrid[y]) - 1; x > 0; x-- {
				if gameGrid[x][y] == 0 {
					continue
				}
				if gameGrid[x-1][y] == 0 {
					gameGrid[x-1][y] = gameGrid[x][y]
					gameGrid[x][y] = 0
					for x2 := x; x2 < len(gameGrid[y])-1; x2++ {
						gameGrid[x2][y] = gameGrid[x2+1][y]
					}
					gameGrid[len(gameGrid[y])-1][y] = 0
				} else if gameGrid[x-1][y] == gameGrid[x][y] {
					gameGrid[x-1][y] *= 2
					for x2 := x; x2 < len(gameGrid[y])-1; x2++ {
						gameGrid[x2][y] = gameGrid[x2+1][y]
					}
					gameGrid[len(gameGrid[y])-1][y] = 0
				}
			}
		}
	}

	i := 0
	for x := range len(gameGrid) {
		for y := range len(gameGrid[x]) {
			if gameGrid[x][y] == 0 {
				i++
			}
		}
	}
	if i <= 0 {
		return gameGrid
	}

	if rl.IsKeyPressed(rl.KeyUp) || rl.IsKeyPressed(rl.KeyDown) || rl.IsKeyPressed(rl.KeyLeft) || rl.IsKeyPressed(rl.KeyRight) {
		x := rand.IntN(len(gameGrid))
		y := rand.IntN(len(gameGrid[0]))
		for !(gameGrid[x][y] == 0) {
			x = rand.IntN(len(gameGrid))
			y = rand.IntN(len(gameGrid[0]))
		}
		gameGrid[x][y] = 2
	}
	return gameGrid

}

func initGrid() [4][4]int {
	var gameGrid [4][4]int
	x := rand.IntN(len(gameGrid))
	y := rand.IntN(len(gameGrid[0]))
	gameGrid[x][y] = 2
	for !(gameGrid[x][y] == 0) {
		x = rand.IntN(len(gameGrid))
		y = rand.IntN(len(gameGrid[0]))
	}
	gameGrid[x][y] = 2
	return gameGrid
}

func main() {
	rl.InitWindow(720, 720, "2048 ig")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	gameGrid := initGrid()

	gameOver := false

	gameWon := false

	var dt float32

	for !rl.WindowShouldClose() {
		dt += rl.GetFrameTime()
		if dt >= 3 {
			dt = 0
			gameOver = emulateInputs(gameGrid)
			for x := range len(gameGrid) {
				for y := range len(gameGrid[x]) {
					if gameGrid[x][y] == 2048 {
						gameWon = true
					}
				}
			}
		}
		if !gameOver && !gameWon {
			gameGrid = runInputs(gameGrid)
		}

		if rl.IsKeyPressed(rl.KeyR) {
			gameGrid = initGrid()
			gameOver = false
			gameWon = false
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.Gray)
		for x := range len(gameGrid) {
			for y := range len(gameGrid[x]) {
				rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.LightGray)
				if gameGrid[x][y] == 0 {
					continue
				}
				if gameGrid[x][y] == 2 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Red)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+79), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 4 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Orange)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+80), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 8 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Yellow)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+80), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 16 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.DarkGreen)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+72), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 32 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Lime)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+64), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 64 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Green)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+64), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 128 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.SkyBlue)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+57), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 256 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Blue)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+52), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 512 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Purple)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+57), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 1024 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Violet)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+42), int32(y*180+68), 50, rl.Black)
				}
				if gameGrid[x][y] == 2048 {
					rl.DrawRectangle(int32(x*180+5), int32(y*180+5), int32(170), int32(170), rl.Pink)
					rl.DrawText(strconv.Itoa(gameGrid[x][y]), int32(x*180+33), int32(y*180+68), 50, rl.Black)
				}
			}
		}
		if gameOver {
			rl.DrawRectangle(80, 300, 580, 100, rl.Black)
			rl.DrawRectangle(115, 400, 520, 60, rl.Black)
			rl.DrawText("Game over!", 100, 305, 100, rl.Magenta)
			rl.DrawText("Press R to restart", 125, 405, 50, rl.Magenta)
		}
		if gameWon {
			rl.DrawRectangle(140, 300, 460, 100, rl.Black)
			rl.DrawRectangle(115, 400, 520, 60, rl.Black)
			rl.DrawText("You won!", 160, 305, 100, rl.Magenta)
			rl.DrawText("Press R to restart", 125, 405, 50, rl.Magenta)
		}
		rl.EndDrawing()
	}
}
