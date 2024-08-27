package main

import (
	"math/rand/v2"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	title  string = "raylib example - basic window"
	width  int32  = 1024
	height int32  = 768
	nBalls int    = 1024
)

type ball struct {
	Position rl.Vector2
	Speed    rl.Vector2
	Radius   float32
}

func main() {
	rl.InitWindow(width, height, title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	balls := NewRandomBallSlice(nBalls)

	for !rl.WindowShouldClose() {
		for i := range balls {
			balls[i].UpdatePosition()
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		rl.DrawFPS(0, 0)

		for i := range balls {
			balls[i].Draw()
		}

		rl.EndDrawing()
	}
}

// Generate n random balls
func NewRandomBallSlice(size int) []ball {
	balls := make([]ball, size)
	for i := 0; i < size; i++ {
		balls[i] = ball{
			Position: rl.Vector2{
				X: rand.Float32() * float32(rl.GetScreenWidth()),
				Y: rand.Float32() * float32(rl.GetScreenHeight()),
			},
			Speed: rl.Vector2{
				X: rand.Float32()*10 - 5, // Random x velocity between -5 and 5
				Y: rand.Float32()*10 - 5, // Random y velocity between -5 and 5
			},
			Radius: rand.Float32()*10 + 10, // Random radius between 10 and 20
		}
	}

	return balls
}

// TODO: balls with low speed vectors are getting "stuck" near the boundary, oscillating back and forth without making meaningful progress
func (b *ball) UpdatePosition() {
	b.Position.X += b.Speed.X
	b.Position.Y += b.Speed.Y

	xMin := b.Radius
	xMax := float32(width) - b.Radius
	if b.Position.X <= xMin || b.Position.X >= xMax {
		b.Speed.X *= float32(-1)
	}

	yMin := b.Radius
	yMax := float32(height) - b.Radius
	if b.Position.Y <= yMin || b.Position.Y >= yMax {
		b.Speed.Y *= float32(-1)
	}
}

func (b *ball) Draw() {
	rl.DrawCircleV(b.Position, b.Radius, rl.Blue)
}
