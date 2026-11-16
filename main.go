package main

import (
	"shakti/src/gui"
	"shakti/src/parsing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(1280, 800, "Shakti")
	defer rl.CloseWindow()

	gui.Initialize()
	defer gui.Close()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		gui.Render()
		parsing.Activate()
	}
}
