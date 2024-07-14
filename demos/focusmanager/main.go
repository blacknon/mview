// Demo code for the FocusManager utility.
package main

import (
	"log"

	"code.rocketnine.space/tslocum/cbind"
	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

func wrap(f func()) func(ev *tcell.EventKey) *tcell.EventKey {
	return func(ev *tcell.EventKey) *tcell.EventKey {
		f()
		return nil
	}
}

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	input1 := mview.NewInputField()
	input1.SetLabel("InputField 1")

	input2 := mview.NewInputField()
	input2.SetLabel("InputField 2")

	input3 := mview.NewInputField()
	input3.SetLabel("InputField 3")

	input4 := mview.NewInputField()
	input4.SetLabel("InputField 4")

	grid := mview.NewGrid()
	grid.SetBorder(true)
	grid.SetTitle(" Press Tab to advance focus ")
	grid.AddItem(input1, 0, 0, 1, 1, 0, 0, true)
	grid.AddItem(input2, 0, 1, 1, 1, 0, 0, false)
	grid.AddItem(input3, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(input4, 1, 0, 1, 1, 0, 0, false)

	focusManager := mview.NewFocusManager(app.SetFocus)
	focusManager.SetWrapAround(true)
	focusManager.Add(input1, input2, input3, input4)

	inputHandler := cbind.NewConfiguration()
	for _, key := range mview.Keys.MovePreviousField {
		err := inputHandler.Set(key, wrap(focusManager.FocusPrevious))
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, key := range mview.Keys.MoveNextField {
		err := inputHandler.Set(key, wrap(focusManager.FocusNext))
		if err != nil {
			log.Fatal(err)
		}
	}
	app.SetInputCapture(inputHandler.Capture)

	app.SetRoot(grid, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
