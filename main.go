package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	height int32 = 450
	width  int32 = 450
	title        = "raylib example - basic window"
)

type ball struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
}

func (b *ball) UpdatePosition() {
	b.Position.X += b.Speed.X
	b.Position.Y += b.Speed.Y

	// x bounds checking
	var (
		xMin                   float32 = b.Radius
		xMax                   float32 = float32(width) - b.Radius
		isXPositionOutOfBounds bool    = b.Position.X <= xMin || b.Position.X >= xMax
	)

	if isXPositionOutOfBounds {
		b.Speed.X *= float32(-1)
	}

	// y bounds checking
	var (
		yMin                   float32 = b.Radius
		yMax                   float32 = float32(height) - b.Radius
		isYPositionOutOfBounds bool    = b.Position.Y <= yMin || b.Position.Y >= yMax
	)

	if isYPositionOutOfBounds {
		b.Speed.Y *= float32(-1)
	}
}

func main() {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	ball := ball{
		Position: rl.Vector2{
			X: float32(width) / float32(2),
			Y: float32(height) / float32(2),
		},
		Speed: rl.Vector2{
			X: 5.0,
			Y: 2.5,
		},
		Radius: float32(20),
	}

	for !rl.WindowShouldClose() {
		// update
		ball.UpdatePosition()

		// draw
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawCircleV(ball.Position, ball.Radius, rl.Blue)
		rl.EndDrawing()
	}
}
