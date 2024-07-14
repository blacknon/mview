// Demo code for the DropDown primitive.
package main

import "github.com/blacknon/mview"

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	dropdown := mview.NewDropDown()
	dropdown.SetLabel("Select an option (hit Enter): ")
	dropdown.SetOptions(nil,
		mview.NewDropDownOption("First"),
		mview.NewDropDownOption("Second"),
		mview.NewDropDownOption("Third"),
		mview.NewDropDownOption("Fourth"),
		mview.NewDropDownOption("Fifth"))

	app.SetRoot(dropdown, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
