package main

import (
	ray "github.com/gen2brain/raylib-go/raylib"
	textbox "raytodo"
)

func main() {
	textbox1 := textbox.NewTextbox(100, 100, 600, 50, 32)
	notes := make([]string, 0)

	ray.InitWindow(840, 600, "Quick-Task")
	ray.SetExitKey(ray.KeyQ)
	defer ray.CloseWindow()
	ray.SetTargetFPS(60)
	for !ray.WindowShouldClose() {
		ray.BeginDrawing()
		ray.ClearBackground(ray.Black)
		textbox1.HandleInput()
		textbox1.Render()
		var but ray.Rectangle
		but = ray.NewRectangle(725, 100, 80, 45)
		ray.DrawRectangleRec(but, ray.LightGray)
		ray.DrawText("Add", 745, 110, 20, ray.Black)
		if ray.IsMouseButtonPressed(ray.MouseLeftButton) {
			if ray.CheckCollisionPointRec(ray.GetMousePosition(), but) {
				notes = append(notes, textbox1.Text)
				textbox1.Text = ""
			}
		}
		ray.DrawText("Quick Task List", 100, 50, 20, ray.LightGray)
		if ray.IsKeyPressed(ray.KeyEnter) {

			notes = append(notes, textbox1.Text)
			textbox1.Text = ""
		}
		yPosition := 200
		for _, i := range notes {
			ray.DrawText(i, 100, int32(yPosition), 20, ray.Green)
			rbut := ray.NewRectangle(600, float32(yPosition), 100, 40)
			ray.DrawRectangleRec(rbut, ray.LightGray)
			ray.DrawText("Remove", 610, int32(yPosition), 20, ray.Black)
			if ray.IsMouseButtonPressed(ray.MouseLeftButton) {
				if ray.CheckCollisionPointRec(ray.GetMousePosition(), rbut) {
          notes = notes[:len(notes)-1]
					textbox1.Text = ""
				}
			}
			yPosition += 50
		}
		ray.DrawText("Made using Raylib", 650, 550, 18, ray.LightGray)
		ray.EndDrawing()
	}

}
