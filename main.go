package main

import (
	"fmt"
	"math/rand"
	"os"
)

// Change game parameters here
const height int = 10
const width int = 10
const numberOfMines int = 10

//

var opened int = 0
var totalFields int = 0

func main() {
	var realBoard [][]int = createBoard(height, width)
	fillWithMines(realBoard, numberOfMines)
	// printBoard(realBoard)
	enrichBoard(realBoard)
	var userBoard [][]int = createBoard(height, width)
	// println()
	// printBoard(realBoard)
	fmt.Println("User board:")
	printUserBoard(userBoard)

	for {
		fmt.Printf("Remaining fields: %d\n", totalFields-opened-numberOfMines)
		var selectedRow int
		var selectedCol int
		fmt.Println("Select a row: ")
		_, _ = fmt.Scanf("%d", &selectedRow)
		fmt.Println("Select a column: ")
		_, _ = fmt.Scanf("%d", &selectedCol)
		clickField(realBoard, userBoard, selectedRow, selectedCol)
		printUserBoard(userBoard)
		if totalFields == (numberOfMines + opened) {
			fmt.Println("You won!")
			os.Exit(0)
		}
	}
}

func createBoard(height int, width int) [][]int {
	totalFields = height * width
	board := make([][]int, height)
	for i := 0; i < height; i++ {
		board[i] = make([]int, width)
	}
	return board
}

func fillWithMines(board [][]int, numberOfMines int) {
	for i := 0; i < numberOfMines; i++ {
		for {
			var field int = board[rand.Intn(len(board))][rand.Intn(len(board[0]))]
			if field != -1 {
				board[rand.Intn(len(board))][rand.Intn(len(board[0]))] = -1
				break
			}
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
			} else {
				fmt.Printf("%d ", board[i][j])
			}
		}
		println()
	}
}

func printUserBoard(board [][]int) {
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

func enrichBoard(board [][]int) {
	// find field with value -1 and increment by 1 each surrounding field
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == -1 {

				// Look at N up
				if i > 0 {
					if board[i-1][j] != -1 {
						board[i-1][j]++
					}
				}

				// Look at NE up right
				if i > 0 && j < len(board[0])-1 {
					if board[i-1][j+1] != -1 {
						board[i-1][j+1]++
					}
				}

				// Look at E right
				if j < len(board[0])-1 {
					if board[i][j+1] != -1 {
						board[i][j+1]++
					}
				}

				// Look at SE right down
				if i < len(board)-1 && j < len(board[0])-1 {
					if board[i+1][j+1] != -1 {
						board[i+1][j+1]++
					}
				}

				// Look at S down
				if i < len(board)-1 {
					if board[i+1][j] != -1 {
						board[i+1][j]++
					}
				}

				// Look at SW left down
				if i < len(board)-1 && j > 0 {
					if board[i+1][j-1] != -1 {
						board[i+1][j-1]++
					}
				}

				// Look at W left
				if j > 0 {
					if board[i][j-1] != -1 {
						board[i][j-1]++
					}
				}

				// Look at NW left up
				if i > 0 && j > 0 {
					if board[i-1][j-1] != -1 {
						board[i-1][j-1]++
					}
				}
			}
		}
	}
}

func clickField(realBoard [][]int, userBoard [][]int, row int, col int) {
	if userBoard[row][col] == -3 || userBoard[row][col] > 0 {
		// fmt.Println("You can not click on already uncovered field!")
		return
	}

	opened++
	if realBoard[row][col] == -1 {
		fmt.Println("BOOOOOOOOM!")
		printBoard(realBoard)
		os.Exit(0)
	}

	if realBoard[row][col] == 0 {
		userBoard[row][col] = -3

		//Find surrounding fields and if it's not a mine then click it
		// Look at N up
		if row > 0 {
			if realBoard[row-1][col] != -1 {
				clickField(realBoard, userBoard, row-1, col)
			}
		}

		// Look at NE up right
		if row > 0 && col < len(realBoard[0])-1 {
			if realBoard[row-1][col+1] != -1 {
				clickField(realBoard, userBoard, row-1, col+1)
			}
		}

		// Look at E right
		if col < len(realBoard[0])-1 {
			if realBoard[row][col+1] != -1 {
				clickField(realBoard, userBoard, row, col+1)
			}
		}

		// Look at SE right down
		if row < len(realBoard)-1 && col < len(realBoard[0])-1 {
			if realBoard[row+1][col+1] != -1 {
				clickField(realBoard, userBoard, row+1, col+1)
			}
		}

		// Look at S down
		if row < len(realBoard)-1 {
			if realBoard[row+1][col] != -1 {
				clickField(realBoard, userBoard, row+1, col)
			}
		}

		// Look at SW left down
		if row < len(realBoard)-1 && col > 0 {
			if realBoard[row+1][col-1] != -1 {
				clickField(realBoard, userBoard, row+1, col-1)
			}
		}

		// Look at W left
		if col > 0 {
			if realBoard[row][col-1] != -1 {
				clickField(realBoard, userBoard, row, col-1)
			}
		}

		// Look at NW left up
		if row > 0 && col > 0 {
			if realBoard[row-1][col-1] != -1 {
				clickField(realBoard, userBoard, row-1, col-1)
			}
		}
		return
	}

	userBoard[row][col] = realBoard[row][col]
}
