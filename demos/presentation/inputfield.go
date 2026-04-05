package main

import (
	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v3"
)

// InputField demonstrates the InputField.
func InputField(nextSlide func()) (title string, info string, content mview.Primitive) {
	input := mview.NewInputField()
	input.SetLabel("Enter a number: ")
	input.SetAcceptanceFunc(mview.InputFieldInteger)
	input.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "InputField", "", Code(input, 30, 1, "inputfield")
}
