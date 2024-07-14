package main

import (
	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

func demoBox(title string) *mview.Box {
	b := mview.NewBox()
	b.SetBorder(true)
	b.SetTitle(title)
	return b
}

// Flex demonstrates flexbox layout.
func Flex(nextSlide func()) (title string, info string, content mview.Primitive) {
	modalShown := false
	panels := mview.NewPanels()

	textView := mview.NewTextView()
	textView.SetBorder(true)
	textView.SetTitle("Flexible width, twice of middle column")
	textView.SetDoneFunc(func(key tcell.Key) {
		if modalShown {
			nextSlide()
			modalShown = false
		} else {
			panels.ShowPanel("modal")
			modalShown = true
		}
	})

	subFlex := mview.NewFlex()
	subFlex.SetDirection(mview.FlexRow)
	subFlex.AddItem(demoBox("Flexible width"), 0, 1, false)
	subFlex.AddItem(demoBox("Fixed height"), 15, 1, false)
	subFlex.AddItem(demoBox("Flexible height"), 0, 1, false)

	flex := mview.NewFlex()
	flex.AddItem(textView, 0, 2, true)
	flex.AddItem(subFlex, 0, 1, false)
	flex.AddItem(demoBox("Fixed width"), 30, 1, false)

	modal := mview.NewModal()
	modal.SetText("Resize the window to see the effect of the flexbox parameters")
	modal.AddButtons([]string{"Ok"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		panels.HidePanel("modal")
	})

	panels.AddPanel("flex", flex, true, true)
	panels.AddPanel("modal", modal, false, false)
	return "Flex", "", panels
}
