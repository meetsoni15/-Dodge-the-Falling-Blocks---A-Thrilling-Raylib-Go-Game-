package main

import (
	_ "embed"
	"fmt"
	_ "image/png" // Required for loading PNG images

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	// Game settings
	const windowWidth, windowHeight = 800, 600

	rl.InitWindow(windowWidth, windowHeight, "Dodge the Falling Blocks")

	// Load embedded textures
	background := loadTextureFromImage(backgroundImage)
	defer rl.UnloadTexture(background)

	// Display splash screen
	fpsSelection := true
	fpsOptions := []int{30, 60, 90, 120, 240}
	selectedFPS := 60

	for fpsSelection {
		rl.BeginDrawing()
		rl.DrawTexture(background, 0, 0, rl.White)

		rl.DrawText("Select FPS", windowWidth/2-50, windowHeight/2-100, 30, rl.White)

		for i, fps := range fpsOptions {
			yPos := windowHeight/2 + i*30
			text := fmt.Sprintf("[%d FPS] %s", fps, ifElse(fps == selectedFPS, "<- Selected", ""))
			rl.DrawText(text, windowWidth/2-100, int32(yPos), 20, rl.White)
		}

		rl.DrawText("Press UP/DOWN to navigate, ENTER to confirm", windowWidth/2-200, windowHeight-50, 20, rl.White)
		rl.EndDrawing()

		if rl.IsKeyPressed(rl.KeyDown) {
			selectedFPS = nextFPS(fpsOptions, selectedFPS)
		}
		if rl.IsKeyPressed(rl.KeyUp) {
			selectedFPS = previousFPS(fpsOptions, selectedFPS)
		}
		if rl.IsKeyPressed(rl.KeyEnter) {
			fpsSelection = false
		}
	}

	rl.SetTargetFPS(int32(selectedFPS))
	rl.CloseWindow()
}
