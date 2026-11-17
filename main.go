package main

import (
	"shakti/src/gui"
	"shakti/src/parsing"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	urlChannel := make(chan string)

	rl.SetConfigFlags(rl.FlagWindowHighdpi)
	rl.InitWindow(1280, 800, "Shakti")
	defer rl.CloseWindow()

	gui.Initialize()
	defer gui.Close()

	rl.SetTargetFPS(60)

	go parsing.Activate(urlChannel)

	for !rl.WindowShouldClose() {
		gui.Render(urlChannel)
	}
}
