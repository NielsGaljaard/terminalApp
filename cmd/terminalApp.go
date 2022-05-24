package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	form := tview.NewInputField().SetAcceptanceFunc(tview.InputFieldInteger).
		SetFieldWidth(100).
		SetLabel("enter number")
	form.SetDoneFunc(func(key tcell.Key) {
		input := form.GetText()
		form.SetText("")
		app.Stop()
		fmt.Println(input)
	})
	box := tview.NewBox().SetBorder(true).SetTitle("hello world")
	grid := tview.NewGrid()
	grid.AddItem(form, 0, 0, 1, 1, 0, 0, false)
	grid.AddItem(box, 1, 0, 9, 1, 0, 0, false)
	grid.SetBorder(true).SetTitle("box")
	if err := app.SetRoot(grid, true).SetFocus(form).Run(); err != nil {
		panic(err)
	}
}
