// Demo code for the Frame primitive.
package main

import (
	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	box := mview.NewBox()
	box.SetBackgroundColor(tcell.ColorBlue.TrueColor())

	frame := mview.NewFrame(box)
	frame.SetBorders(2, 2, 2, 2, 4, 4)
	frame.AddText("Header left", true, mview.AlignLeft, tcell.ColorWhite.TrueColor())
	frame.AddText("Header middle", true, mview.AlignCenter, tcell.ColorWhite.TrueColor())
	frame.AddText("Header right", true, mview.AlignRight, tcell.ColorWhite.TrueColor())
	frame.AddText("Header second middle", true, mview.AlignCenter, tcell.ColorRed.TrueColor())
	frame.AddText("Footer middle", false, mview.AlignCenter, tcell.ColorGreen.TrueColor())
	frame.AddText("Footer second middle", false, mview.AlignCenter, tcell.ColorGreen.TrueColor())

	app.SetRoot(frame, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
