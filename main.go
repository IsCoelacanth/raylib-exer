package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
	"math/rand/v2"
)

type Square struct {
	x, y, h, w int32
	dx, dy     float32
}

type WindowSpec struct {
	h, w, fps int32
	title     string
	win_color rl.Color
}

func rand_velocity() float64 {
	velo := rand.Float64()*10 - 5.0

	// ensure at least 1 unit / tick
	if velo > 0 {
		velo = max(2, velo)
	} else {
		velo = min(-2, velo)
	}
	return velo
}

func random_color() rl.Color {
	return rl.ColorFromHSV(float32(rand.IntN(360)), rand.Float32(), rand.Float32())
}

func (bx *Square) animate(win WindowSpec) WindowSpec {
	var flip_window_color bool = false
	var both int = 0
	if bx.x <= 0 || bx.x+bx.w >= win.w {
		bx.dx *= -1
		flip_window_color = true
		both++
	}
	if bx.y <= 0 || bx.y+bx.h >= win.h {
		bx.dy *= -1
		flip_window_color = true
		both++
	}
	if flip_window_color {
		win.win_color = random_color()
	}
	if both == 2 {
		rl.DrawText("CORNER!!!!", win.w/2-40, win.h/2-20, 20, rl.LightGray)
	}
	bx.x = int32(float32(bx.x) + bx.dx + 0.5)
	bx.y = int32(float32(bx.y) + bx.dy + 0.5)
	rl.DrawRectangle(bx.x, bx.y, bx.h, bx.w, rl.RayWhite)
	return win
}

func main() {
	var box Square = Square{x: int32(rand.IntN(760)), y: int32(rand.IntN(560)), h: 20, w: 20, dx: float32(rand_velocity()), dy: float32(rand_velocity())}
	win_specs := WindowSpec{h: 600, w: 800, fps: 120, title: "Test Window", win_color: random_color()}

	fmt.Println("Starting Service", win_specs, box)
	// takes w, and h
	rl.InitWindow(win_specs.w, win_specs.h, win_specs.title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(win_specs.fps)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(win_specs.win_color)
		win_specs = box.animate(win_specs)
		// fmt.Println(box, win_specs)
		rl.EndDrawing()
	}
}
