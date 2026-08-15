package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// declare constants
const (
	screenWidth  = 800
	screenHeight = 450
)

// declare global variables
var (
	running = true
)

func Init() {
	rl.InitWindow(screenWidth, screenHeight, "GoBreakBricks")

	rl.SetTargetFPS(60)
}

// func Input(g *Game) {
//
// }

func Update() {
}

func Draw() {
	// Draw
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	rl.DrawText("Welcome to GoBreakBricks", 0, 0, 24, rl.White)

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
