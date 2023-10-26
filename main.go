package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	var isGameRunning bool // Indicates if the game is running

	mainText := tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText("Main content (Game Not Running)")

	updateMainText := func() {
		text := "Main content"
		if isGameRunning {
			text += " (Game Running)"
		} else {
			text += " (Game Not Running)"
		}
		mainText.SetText(text)
	}

	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}

	app := tview.NewApplication()

	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(-1/3, 0).
		SetBorders(true).
		SetBordersColor(tcell.Color169).
		AddItem(newPrimitive("Super Tic-Tac-Toe"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("topher-nullset ; made with love"), 2, 0, 1, 3, 0, 0, false)

	var list *tview.List
	var endGame func()

	newGame := func() {
		if isGameRunning {
			return
		}
		isGameRunning = true
		updateMainText()
		list.AddItem("Press M to return to Menu", "", 'M', nil)
	}

	endGame = func() {
		isGameRunning = false
		updateMainText()
		list.RemoveItem(4)
		app.SetRoot(grid, true).SetFocus(list)
	}

	quitConfirm := func() {
		modal := tview.NewModal().
			SetText("Are you sure you want to quit?").
			AddButtons([]string{"Yes", "No"}).
			SetBackgroundColor(tcell.ColorBlack).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				if buttonLabel == "Yes" {
					app.Stop()
				} else {
					app.SetRoot(grid, true).SetFocus(list)
				}
			})
		app.SetRoot(modal, false).SetFocus(modal)
	}

	difficulty := func() {
		if isGameRunning {
			return
		}
		// Placeholder for Difficulty logic
	}

	statistics := func() {
		if isGameRunning {
			return
		}
		// Placeholder for Statistics logic
	}

	list = tview.NewList().
		AddItem("New Game", "", '1', newGame).
		AddItem("Difficulty", "", '2', difficulty).
		AddItem("Statistics", "", '3', statistics).
		AddItem("Quit", "", '4', quitConfirm)

	menu := tview.NewFlex().
		AddItem(list, 0, 1, true)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(mainText, 1, 1, 1, 2, 0, 0, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if isGameRunning {
			if event.Rune() == 'M' || event.Rune() == 'm' {
				endGame()
				return nil
			}
			return nil
		}
		return event
	})

	if err := app.SetRoot(grid, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
