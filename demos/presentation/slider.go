package main

import (
	"fmt"

	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v3"
)

// Slider demonstrates the Slider.
func Slider(nextSlide func()) (title string, info string, content mview.Primitive) {
	slider := mview.NewSlider()
	slider.SetLabel("Volume:   0%")
	slider.SetChangedFunc(func(value int) {
		slider.SetLabel(fmt.Sprintf("Volume: %3d%%", value))
	})
	slider.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "Slider", sliderInfo, Code(slider, 30, 1, "slider")
}
