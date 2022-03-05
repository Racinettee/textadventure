package textadventure

import (
	ui "github.com/VladimirMarkelov/clui"
)

type TextWindow struct {
	*ui.Window
	// The object managing the text we view
	Buffer Buffer
	// The line from which we start rendering
	TopLine int
}

func CreateTextWindow(x, y, w, h int, title string) *TextWindow {
	textWin := new(TextWindow)
	textWin.Window = ui.AddWindow(x, y, w, h, title)
	return textWin
}
