// Demo code for the Flex primitive.
package main

import (
	"github.com/blacknon/mview"
)

func demoBox(title string) *mview.Box {
	b := mview.NewBox()
	b.SetBorder(true)
	b.SetTitle(title)
	return b
}

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	app.EnableMouse(true)

	subFlex := mview.NewFlex()
	subFlex.SetDirection(mview.FlexRow)
	subFlex.AddItem(demoBox("Top"), 0, 1, false)
	subFlex.AddItem(demoBox("Middle (3 x height of Top)"), 0, 3, false)
	subFlex.AddItem(demoBox("Bottom (5 rows)"), 5, 1, false)

	flex := mview.NewFlex()
	flex.AddItem(demoBox("Left (1/2 x width of Top)"), 0, 1, false)
	flex.AddItem(subFlex, 0, 2, false)
	flex.AddItem(demoBox("Right (20 cols)"), 20, 1, false)

	app.SetRoot(flex, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
