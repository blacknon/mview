// Demo code for the CheckBox primitive.
package main

import (
	"github.com/blacknon/mview"
)

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	checkbox := mview.NewCheckBox()
	checkbox.SetLabel("Hit Enter to check box: ")

	app.SetRoot(checkbox, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
