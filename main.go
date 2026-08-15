package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// declare constants
const (
	screenWidth  = 800
	screenHeight = 450
)

// declare global variables
var (
	running  = true
	showMenu = true
)

func Init() {
	rl.InitWindow(
		screenWidth,
		screenHeight,
		"GoBreakBricks",
	)

	rl.SetTargetFPS(60)
}

func ShowMenu() {
	text := "Welcome to GoBreakBricks"
	var fontSize int32 = 24
	textWidth := int32(rl.MeasureText(text, fontSize))

	var btnWidth float32 = 200
	var btnHeight float32 = 80
	btnBounds := rl.NewRectangle(
		(screenWidth-btnWidth)/2,
		(screenHeight-btnHeight)/2,
		btnWidth,
		btnHeight,
	)

	rl.DrawText(
		text,
		(screenWidth-textWidth)/2,
		0,
		24,
		rl.White,
	)

	if gui.Button(btnBounds, "Start Game") {
		showMenu = false
	}
}

// func Input(g *Game) {
//
// }

func Update() {
}

func Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	if showMenu {
		ShowMenu()
	}

	rl.EndDrawing()
}

func main() {
	Init()

	running = !rl.WindowShouldClose()

	for running {

		// Input()

		Update()

		Draw()

		// handle window close
		running = !rl.WindowShouldClose()
	}
}
