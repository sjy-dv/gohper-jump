package main

import (
	"math/rand"
	"strconv"
	"time"

	rl "github.com/lachee/raylib-goplus/raylib"
	// rl "github.com/gen2brain/raylib-go/raylib"
)

type PlatForm struct {
	Position_X int
	Position_Y int
	Color      rl.Color
}

func main() {
	var (
		screenWidth  int = 450
		screenHeight int = 800
	)

	rl.InitWindow(screenWidth, screenHeight, "gohper jump")
	rl.SetTargetFPS(60)

	backGroundSource := rl.LoadImage("assets/background.png")
	rl.ImageResize(backGroundSource, screenWidth, screenHeight)

	backGround := rl.LoadTextureFromImage(backGroundSource)

	gohperSource := rl.LoadImage("assets/gohper.png")

	gohper := rl.LoadTextureFromImage(gohperSource)
	var x_pos int = 0
	var y_pos int = 0
	// you can change this

	platforms := []PlatForm{}
	platform1 := PlatForm{10, 100, rl.Black}
	platform2 := PlatForm{100, 200, rl.Black}
	platform3 := PlatForm{300, 300, rl.Brown}
	platform4 := PlatForm{40, 400, rl.Black}
	platform5 := PlatForm{250, 500, rl.Black}
	platform6 := PlatForm{450, 600, rl.Brown}
	platform7 := PlatForm{450, 700, rl.Black}
	platform8 := PlatForm{250, 750, rl.Black}

	platforms = append(platforms, platform1)
	platforms = append(platforms, platform2)
	platforms = append(platforms, platform3)
	platforms = append(platforms, platform4)
	platforms = append(platforms, platform5)
	platforms = append(platforms, platform6)
	platforms = append(platforms, platform7)
	platforms = append(platforms, platform8)

	rand.Seed(time.Now().UnixNano())

	var score int = 0
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.DrawTexture(backGround, 0, 0, rl.White)

		rl.DrawTexture(gohper, x_pos, y_pos, rl.White)
		rl.DrawText("Score: "+strconv.Itoa(score), 0, 0, 20, rl.Black)
		for index, current_platform := range platforms {
			rl.DrawRectangle(current_platform.Position_X, current_platform.Position_Y, 100, 30, current_platform.Color)

			if rl.CheckCollisionRecs(rl.NewRectangle(float32(x_pos), float32(y_pos), float32(52), float32(120)), rl.NewRectangle(float32(current_platform.Position_X-37), float32(current_platform.Position_Y), float32(100), float32(30))) {
				y_pos -= 120
				score++
				if current_platform.Color == rl.Brown {
					var PosX int = (rand.Intn(350))
					platforms[index].Position_X = PosX
					platforms[index].Position_Y = 0
				}

			}
			platforms[index].Position_Y += 1
			if current_platform.Position_Y > screenHeight {
				var PosX int = (rand.Intn(350))
				platforms[index].Position_X = PosX
				platforms[index].Position_Y = 0
			}
		}

		if rl.IsKeyDown(rl.KeyA) && x_pos < 400 && x_pos > -30 {
			x_pos -= 5
		}
		if rl.IsKeyDown(rl.KeyD) && x_pos < 400 && x_pos > -30 {
			x_pos += 5
		}
		if x_pos >= 400 {
			x_pos -= 5
		}
		if x_pos < 0 {
			x_pos += 5
		}

		if y_pos > screenHeight {
			platforms = nil
			rl.UnloadTexture(gohper)
			rl.ClearBackground(rl.White)
			rl.DrawText("Your final score is: "+strconv.Itoa(score), 30, 40, 30, rl.Black)
		}
		rl.EndDrawing()
		y_pos += 2
	}
	rl.CloseWindow()
}