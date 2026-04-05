package main

import (
	"github.com/blacknon/mview"
)

// The width of the code window.
const codeWidth = 56

// Code returns a primitive which displays the given primitive (with the given
// size) on the left side and its source code on the right side.
func Code(p mview.Primitive, width, height int, name string) mview.Primitive {
	// Set up code view.
	codeView := mview.NewTextView()
	codeView.SetWrap(true)
	codeView.SetWordWrap(true)
	codeView.SetDynamicColors(false)
	codeView.SetPadding(1, 1, 2, 0)
	codeView.Write(exampleCode(name))

	f := mview.NewFlex()
	f.AddItem(Center(width, height, p), 0, 1, true)
	f.AddItem(codeView, codeWidth, 1, false)
	return f
}
