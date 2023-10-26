package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Game struct {
	board         [3][3]string
	currentPlayer string
	cells         [3][3]*tview.TextView // Stores the TextView cells for the game board
}

func updateUIWithMessage(message string, textView *tview.TextView) {
	// Add code to update the UI with the provided message
	// For example, you can update the provided TextView widget with the message
	textView.SetText(message)
}

func updateBoardUI(board [3][3]string, cells [3][3]*tview.TextView) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cells[i][j].SetText(board[i][j])
		}
	}
}

func newPrimitive(text string) tview.Primitive {
	return tview.NewTextView().
		SetTextAlign(tview.AlignCenter).
		SetText(text)
}

func newGame() *Game {
	return &Game{currentPlayer: "X"}
}

func endGame(isGameRunning *bool, list *tview.List, app *tview.Application, mainGrid *tview.Grid) {
	*isGameRunning = false
	list.RemoveItem(4)
	app.SetRoot(mainGrid, true).SetFocus(list)
}

func TicTacToeGrid(game *Game) *tview.Grid {
	grid := tview.NewGrid().
		SetRows(3, 3, 3).
		SetColumns(0, 0, 0).
		SetBorders(true).
		SetBordersColor(tcell.ColorDefault)

	// Create the Tic-Tac-Toe grid cells and add them to the grid
	// For simplicity, using TextView for cells, you can use Buttons or other primitives as needed
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			cell := tview.NewTextView().
				SetText("▒▒▒\n▒▒▒\n▒▒▒").
				SetTextAlign(tview.AlignCenter)
			grid.AddItem(cell, i, j, 1, 1, 0, 0, false)
		}
	}

	// Initialize game board and current player
	game.board = [3][3]string{}
	game.currentPlayer = "X"

	return grid
}

func checkWinner(player string, board [3][3]string) bool {
	// Check rows
	for i := 0; i < 3; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}

	// Check columns
	for i := 0; i < 3; i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}

	// Check diagonals
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}

	return false
}

func isBoardFull(board [3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == " " {
				return false
			}
		}
	}
	return true
}


func main() {
	var isGameRunning bool // Indicates if the game is running
	var game *Game

	app := tview.NewApplication()

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

	game = newGame()

	mainGrid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(-1/3, 0).
		SetBorders(true).
		SetBordersColor(tcell.Color169).
		AddItem(newPrimitive("Super Tic-Tac-Toe"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("topher-nullset ; made with love"), 2, 0, 1, 3, 0, 0, false)

	ticTacToeGrid := TicTacToeGrid(game)
	mainGrid.AddItem(ticTacToeGrid, 1, 1, 1, 1, 0, 0, false)

	var list *tview.List

	list = tview.NewList().
		AddItem("New Game", "", '1', func() {
			if isGameRunning {
				return
			}
			game = newGame()
			isGameRunning = true
			updateMainText()
			list.AddItem("Press M to return to Menu", "", 'M', nil)
		}).
		// ... (other list items)
		AddItem("Quit", "", '4', func() {
			modal := tview.NewModal().
				SetText("Are you sure you want to quit?").
				AddButtons([]string{"Yes", "No"}).
				SetBackgroundColor(tcell.ColorBlack).
				SetDoneFunc(func(buttonIndex int, buttonLabel string) {
					if buttonLabel == "Yes" {
						app.Stop()
					} else {
						app.SetRoot(mainGrid, true).SetFocus(list)
					}
				})
			app.SetRoot(modal, false).SetFocus(modal)
		})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if isGameRunning {
			if event.Rune() == 'M' || event.Rune() == 'm' {
				endGame(&isGameRunning, list, app, mainGrid)
				return nil
			}
			if event.Key() == tcell.KeyRune {
				input := event.Rune() - '0' // Convert rune to integer
				row, col := (input-1)/3, (input-1)%3
				if row >= 0 && row < 3 && col >= 0 && col < 3 && game.board[row][col] == " " {
					game.board[row][col] = game.currentPlayer

					if checkWinner(game.currentPlayer, game.board) {
						// Handle game over, show winner, etc.
						endGame(&isGameRunning, list, app, mainGrid)
						return nil
					}

					if game.currentPlayer == "X" {
						game.currentPlayer = "O"
					} else {
						game.currentPlayer = "X"
					}

					updateBoardUI(game.board, game.cells)

					if isBoardFull(game.board) && !checkWinner(game.currentPlayer, game.board) {
						drawMessage := "It's a draw!"
						updateUIWithMessage(drawMessage, mainText)
						endGame(&isGameRunning, list, app, mainGrid)
						return nil
					}
				}
			}
			return nil
		}
		return event
	})

	if err := app.SetRoot(mainGrid, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
