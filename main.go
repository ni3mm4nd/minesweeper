package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"samoseto.com/minesweeper/internal/game"
)

//go:embed templates
var templates embed.FS

func main() {

	port := flag.Int("port", 8081, "port number to listen on")
	flag.Parse()

	router := chi.NewRouter()
	router.Use(middleware.Recoverer)

	router.Get("/", showIndexPage)
	router.Get("/boardclick/{row}/{col}", clickField)
	router.Post("/newgame", newGame)

	fmt.Printf("Listening on localhost:%d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), router)
	catch(err)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

func ParseFiles(filenames ...string) (*template.Template, error) {
	t, err := template.ParseFS(templates, filenames...)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func newGame(w http.ResponseWriter, r *http.Request) {
	rows, _ := strconv.Atoi(r.FormValue("rows"))
	cols, _ := strconv.Atoi(r.FormValue("cols"))
	mines, _ := strconv.Atoi(r.FormValue("mines"))

	game := game.NewGame(rows, cols, mines)

	t, err := ParseFiles("templates/board.html")
	catch(err)

	err = t.Execute(w, *game)
	catch(err)
}

func clickField(w http.ResponseWriter, r *http.Request) {
	row, _ := strconv.Atoi(chi.URLParam(r, "row"))
	col, _ := strconv.Atoi(chi.URLParam(r, "col"))

	game.GetGamePtr().ClickField(row, col)

	t, err := ParseFiles("templates/board.html")
	catch(err)

	err = t.Execute(w, *game.GetGamePtr())
	catch(err)
}

func showIndexPage(w http.ResponseWriter, r *http.Request) {
	data := game.GetGamePtr()
	if data == nil {
		data = game.NewDefaultGame()
	}
	t, err := ParseFiles("templates/layout.html", "templates/index.html", "templates/board.html")
	catch(err)

	err = t.Execute(w, *data)
	catch(err)
}

// For debug purposes
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
