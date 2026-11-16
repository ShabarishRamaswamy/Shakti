package gui

import (
	raygui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const textSize = 24

var (
	font       rl.Font
	searchText string
)

// Initialize loads resources used by the GUI.
func Initialize() {
	scale := rl.GetWindowScaleDPI().X
	font = rl.LoadFontEx("assets/fonts/JetBrainsMono-Regular.ttf", int32(float32(textSize)*scale), nil, 0)
	rl.SetTextureFilter(font.Texture, rl.FilterBilinear)

	raygui.SetFont(font)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_SIZE, textSize)
	raygui.SetStyle(raygui.DEFAULT, raygui.TEXT_COLOR_PRESSED, 0x654321FF)
}

// Close releases GUI resources.
func Close() {
	rl.UnloadFont(font)
}

// Render draws the browser chrome currently implemented by the GUI.
func Render() {
	searchBox := rl.NewRectangle(16, 16, float32(rl.GetScreenWidth()-32), 44)

	rl.BeginDrawing()
	rl.ClearBackground(rl.RayWhite)

	raygui.TextBox(searchBox, &searchText, 1024, true)

	rl.EndDrawing()
}
