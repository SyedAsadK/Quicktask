package textbox

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Textbox struct {
	Rectangle raylib.Rectangle
	Text      string
	MaxLength int
	IsFocused bool
}

func NewTextbox(x, y, width, height float32, maxLength int) *Textbox {
	return &Textbox{
		Rectangle: raylib.NewRectangle(x, y, width, height),
		Text:      "",
		MaxLength: maxLength,
		IsFocused: false,
	}
}

func (tb *Textbox) HandleInput() {
	if raylib.IsMouseButtonPressed(raylib.MouseLeftButton) {
		if raylib.CheckCollisionPointRec(raylib.GetMousePosition(), tb.Rectangle) {
			tb.IsFocused = true
		} else {
			tb.IsFocused = false
		}
	}

	if tb.IsFocused {
		charPressed := int(raylib.GetCharPressed())
		for charPressed > 0 {
			if len(tb.Text) < tb.MaxLength && charPressed >= 32 && charPressed <= 126 {
				tb.Text += string(rune(charPressed))
			}
			charPressed = int(raylib.GetCharPressed())
		}

		if raylib.IsKeyPressed(raylib.KeyBackspace) && len(tb.Text) > 0 {
			tb.Text = tb.Text[:len(tb.Text)-1]
		}
		if raylib.IsKeyDown(raylib.KeyLeftControl) &&
			raylib.IsKeyPressed(raylib.KeyW) {
			tb.Text = tb.Text[:0]
		}
	}
}

func (tb *Textbox) Render() {
	if tb.IsFocused {
		raylib.DrawRectangleRec(tb.Rectangle, raylib.LightGray)
	} else {
		raylib.DrawRectangleRec(tb.Rectangle, raylib.DarkGray)
	}
	raylib.DrawRectangleLinesEx(tb.Rectangle, 2, raylib.Black)

	raylib.DrawText(tb.Text, int32(tb.Rectangle.X+10), int32(tb.Rectangle.Y+10), 20, raylib.Black)
}
