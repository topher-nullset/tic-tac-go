package main

import (
	// "fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)



func main() {
	// app := tview.NewApplication()
	newPrimitive := func(text string) tview.Primitive {
		return tview.NewTextView().
			SetTextAlign(tview.AlignCenter).
			SetText(text)
	}
	menu := newPrimitive("Menu")
	main := newPrimitive("Main content")
	// sideBar := newPrimitive("Side Bar")

	button := tview.NewButton("Button").SetSelectedFunc(func() {
		// app.Stop()
	})
	
	list := tview.NewList().
	AddItem("New Game", "", '1', nil).
	AddItem("Difficulty", "", '2', nil).
	AddItem("Statistics", "", '3', nil)


	grid := tview.NewGrid().
		SetRows(1, 0, 1).
		SetColumns(-1/3, 0).
		SetBorders(true).
		SetBordersColor(tcell.Color169).
		AddItem(newPrimitive("Super Tic-Tac-Toe"), 0, 0, 1, 3, 0, 0, false).
		AddItem(newPrimitive("topher-nullset ; made with love"), 2, 0, 1, 3, 0, 0, false)

		menu = tview.NewFlex().
		AddItem(list, 0, 1, false).
		AddItem(nil, 1, 0, false).
		AddItem(button, 1, 1, false).
		SetDirection(tview.FlexRow)

	grid.AddItem(menu, 1, 0, 1, 1, 0, 0, false).
		AddItem(main, 1, 1, 1, 2, 0, 0, false)
		// AddItem(quitButton, 2, 0, 1, 3, 0, 0, false)
		// AddItem(sideBar, 1, 2, 1, 1, 0, 80, false)

	if err := tview.NewApplication().SetRoot(grid, true).SetFocus(grid).Run(); err != nil {
		panic(err)
	}
}


// var board [3][3]string

// func initBoard() {
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			board[i][j] = " "
// 		}
// 	}
// }

// func clearScreen() {
// 	cmd := exec.Command("clear") // For Linux/macOS
// 	if runtime.GOOS == "windows" {
// 		cmd = exec.Command("cmd", "/c", "cls") // For Windows
// 	}
// 	cmd.Stdout = os.Stdout
// 	cmd.Run()
// }

// func printBoard() {
// 	fmt.Println("  0 1 2")
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("%d %s\n", i, strings.Join(board[i][:], " "))
// 	}
// }

// func isBoardFull() bool {
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if board[i][j] == " " {
// 				return false
// 			}
// 		}
// 	}
// 	return true
// }

// func checkWinner(player string) bool {
// 	// Check rows
// 	for i := 0; i < 3; i++ {
// 		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
// 			return true
// 		}
// 	}

// 	// Check columns
// 	for i := 0; i < 3; i++ {
// 		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
// 			return true
// 		}
// 	}

// 	// Check diagonals
// 	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
// 		return true
// 	}
// 	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
// 		return true
// 	}

// 	return false
// }

// func main() {
// 	initBoard()
// 	currentPlayer := "X"
// 	scanner := bufio.NewScanner(os.Stdin)

// 	for {
// 		clearScreen()
// 		printBoard()
// 		fmt.Printf("Player %s's turn. Enter row and column (e.g., 0 1): ", currentPlayer)
// 		scanner.Scan()
// 		input := scanner.Text()
// 		var row, col int
// 		_, err := fmt.Sscanf(input, "%d %d", &row, &col)
// 		if err != nil || row < 0 || row >= 3 || col < 0 || col >= 3 || board[row][col] != " " {
// 			fmt.Println("Invalid move. Try again.")
// 			continue
// 		}

// 		board[row][col] = currentPlayer

// 		if checkWinner(currentPlayer) {
// 			printBoard()
// 			fmt.Printf("Player %s wins!\n", currentPlayer)
// 			break
// 		}

// 		if isBoardFull() {
// 			clearScreen()
// 			printBoard()
// 			fmt.Println("It's a draw!")
// 			break
// 		}

// 		// Switch players
// 		if currentPlayer == "X" {
// 			currentPlayer = "O"
// 		} else {
// 			currentPlayer = "X"
// 		}
// 	}
// }
