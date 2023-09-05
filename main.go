package main

import (
	"fmt"
	"os"

	"samoseto.com/minesweeper/internal/game"
)

func main() {
	game := game.NewGame(10, 10, 10)
	printBoard(game.UserBoard)

	for {
		fmt.Printf("Remaining fields: %d\n", game.Remaining())
		var selectedRow int
		var selectedCol int
		fmt.Println("Select a row: ")
		_, _ = fmt.Scanf("%d", &selectedRow)
		fmt.Println("Select a column: ")
		_, _ = fmt.Scanf("%d", &selectedCol)
		game.ClickField(selectedRow, selectedCol)
		printBoard(game.UserBoard)
		if game.IsGameOver {
			printBoard(game.RealBoard)
			fmt.Println("Game over!")

			if game.IsWon {
				fmt.Println("You won!")
			} else if game.IsLost {
				fmt.Println("You lost!")
			}

			os.Exit(0)
		}
	}
}

func printBoard(board [][]int) {
	fmt.Printf("  ")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	for i := 0; i < len(board); i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == -1 {
				fmt.Printf("X ")
			} else if board[i][j] == -3 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
		println()
	}
}
