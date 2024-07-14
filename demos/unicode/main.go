// Demo code for unicode support (demonstrates wide Chinese characters).
package main

import (
	"fmt"

	"github.com/blacknon/mview"
)

func main() {
	app := mview.NewApplication()
	defer app.HandlePanic()

	panels := mview.NewPanels()

	form := mview.NewForm()
	form.AddDropDownSimple("称谓", 0, nil, "先生", "女士", "博士", "老师", "师傅")
	form.AddInputField("姓名", "", 20, nil, nil)
	form.AddPasswordField("密码", "", 10, '*', nil)
	form.AddCheckBox("", "年龄 18+", false, nil)
	form.AddButton("保存", func() {
		_, option := form.GetFormItem(0).(*mview.DropDown).GetCurrentOption()
		userName := form.GetFormItem(1).(*mview.InputField).GetText()

		alert(panels, "alert-dialog", fmt.Sprintf("保存成功，%s %s！", userName, option.GetText()))
	})
	form.AddButton("退出", func() {
		app.Stop()
	})
	form.SetBorder(true)
	form.SetTitle("输入一些内容")
	form.SetTitleAlign(mview.AlignLeft)
	panels.AddPanel("base", form, true, true)

	app.SetRoot(panels, true)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

// alert shows a confirmation dialog.
func alert(panels *mview.Panels, id string, message string) {
	modal := mview.NewModal()
	modal.SetText(message)
	modal.AddButtons([]string{"确定"})
	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		panels.HidePanel(id)
		panels.RemovePanel(id)
	})

	panels.AddPanel(id, modal, false, true)
}
