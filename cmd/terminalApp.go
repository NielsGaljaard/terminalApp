package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewInputField().
		SetAcceptanceFunc(tview.InputFieldMaxLength(26)).
		SetFieldWidth(100).
		SetLabel("Query")
	form.SetBackgroundColor(tcell.NewHexColor(0))
	form.SetFieldBackgroundColor(tcell.NewHexColor(0))
	form.SetDoneFunc(func(key tcell.Key) {
		input := form.GetText()
		form.SetText("")
		app.Stop()
		fmt.Println(input)
	}).SetBackgroundColor(tcell.NewHexColor(0))
	formBox := tview.NewFlex()
	formBox.AddItem(form, 20, 1, false)
	formBox.SetDirection(tview.FlexRow).SetBackgroundColor(tcell.NewHexColor(0))
	box := tview.NewBox().SetBackgroundColor(tcell.NewHexColor(0))
	grid := tview.NewGrid()
	grid.SetBackgroundColor(tcell.NewHexColor(0))
	grid.AddItem(formBox, 0, 0, 1, 5, 0, 0, false)
	grid.AddItem(box, 1, 0, 13, 5, 0, 0, false)
	grid.SetBorder(true).SetTitle("box")
	if err := app.SetRoot(grid, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
