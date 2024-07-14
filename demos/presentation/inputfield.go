package main

import (
	"github.com/blacknon/mview"
	"github.com/gdamore/tcell/v2"
)

const inputField = `[green]package[white] main

[green]import[white] (
    [red]"strconv"[white]

    [red]"github.com/gdamore/tcell/v2"[white]
    [red]"github.com/blacknon/mview"[white]
)

[green]func[white] [yellow]main[white]() {
    input := mview.[yellow]NewInputField[white]().
        [yellow]SetLabel[white]([red]"Enter a number: "[white]).
        [yellow]SetAcceptanceFunc[white](
            mview.InputFieldInteger,
        ).[yellow]SetDoneFunc[white]([yellow]func[white](key tcell.Key) {
            text := input.[yellow]GetText[white]()
            n, _ := strconv.[yellow]Atoi[white](text)
            [blue]// We have a number.[white]
        })
    mview.[yellow]NewApplication[white]().
        [yellow]SetRoot[white](input, true).
        [yellow]Run[white]()
}`

// InputField demonstrates the InputField.
func InputField(nextSlide func()) (title string, info string, content mview.Primitive) {
	input := mview.NewInputField()
	input.SetLabel("Enter a number: ")
	input.SetAcceptanceFunc(mview.InputFieldInteger)
	input.SetDoneFunc(func(key tcell.Key) {
		nextSlide()
	})
	return "InputField", "", Code(input, 30, 1, inputField)
}
