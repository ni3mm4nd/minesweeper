package game

import "math/rand"

var gamePtr *gameStruct

type gameStruct struct {
	UserBoard     [][]int
	RealBoard     [][]int
	Height        int
	Width         int
	NumberOfMines int
	Opened        int
	TotalFields   int
	IsGameOver    bool
	IsWon         bool
	IsLost        bool
	Remaining     func() int
}

func GetGamePtr() *gameStruct {
	return gamePtr
}

func NewGame(height int, width int, numberOfMines int) *gameStruct {
	gamePtr = &gameStruct{
		UserBoard:     [][]int{},
		RealBoard:     [][]int{},
		Height:        height,
		Width:         width,
		NumberOfMines: numberOfMines,
		Opened:        0,
		TotalFields:   height * width,
		IsGameOver:    false,
		IsWon:         false,
		IsLost:        false,
		Remaining:     func() int { return gamePtr.TotalFields - gamePtr.Opened - gamePtr.NumberOfMines },
	}

	gamePtr.TotalFields = height * width
	gamePtr.RealBoard = createBoard(height, width)
	fillWithMines(gamePtr.RealBoard, numberOfMines)
	enrichBoard(gamePtr.RealBoard)
	gamePtr.UserBoard = createBoard(height, width)

	return gamePtr
}

func (g *gameStruct) ClickField(row int, col int) {
	if g.IsGameOver {
		return
	}

	if g.UserBoard[row][col] == -3 || g.UserBoard[row][col] > 0 {
		// fmt.Println("You can not click on already uncovered field!")
		return
	}

	g.Opened++
	if g.RealBoard[row][col] == -1 {
		g.IsGameOver = true
		g.IsLost = true
		return
	}

	if g.RealBoard[row][col] == 0 {
		g.UserBoard[row][col] = -3

		//Find surrounding fields and if it's not a mine then click it
		// Look at N up
		if row > 0 {
			if g.RealBoard[row-1][col] != -1 {
				g.ClickField(row-1, col)
			}
		}

		// Look at NE up right
		if row > 0 && col < g.Width-1 {
			if g.RealBoard[row-1][col+1] != -1 {
				g.ClickField(row-1, col+1)
			}
		}

		// Look at E right
		if col < g.Width-1 {
			if g.RealBoard[row][col+1] != -1 {
				g.ClickField(row, col+1)
			}
		}

		// Look at SE right down
		if row < g.Height-1 && col < g.Width-1 {
			if g.RealBoard[row+1][col+1] != -1 {
				g.ClickField(row+1, col+1)
			}
		}

		// Look at S down
		if row < g.Height-1 {
			if g.RealBoard[row+1][col] != -1 {
				g.ClickField(row+1, col)
			}
		}

		// Look at SW left down
		if row < g.Height-1 && col > 0 {
			if g.RealBoard[row+1][col-1] != -1 {
				g.ClickField(row+1, col-1)
			}
		}

		// Look at W left
		if col > 0 {
			if g.RealBoard[row][col-1] != -1 {
				g.ClickField(row, col-1)
			}
		}

		// Look at NW left up
		if row > 0 && col > 0 {
			if g.RealBoard[row-1][col-1] != -1 {
				g.ClickField(row-1, col-1)
			}
		}

		return
	}

	g.UserBoard[row][col] = g.RealBoard[row][col]

	if g.TotalFields == (g.NumberOfMines + g.Opened) {
		g.IsGameOver = true
		g.IsWon = true
		return
	}
}

func enrichBoard(board [][]int) {
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

func createBoard(height int, width int) [][]int {
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
