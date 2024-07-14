package main

import (
	"fmt"

	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

// End shows the final slide.
func End(nextSlide func()) (title string, info string, content mview.Primitive) {
	textView := mview.NewTextView()
	textView.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	url := "https://github.com/blacknon/mview"
	fmt.Fprint(textView, url)
	return "End", "", Center(len(url), 1, textView)
}
