package game

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
)

const marker = "O"

// Empty board spaces represented by white color
var (
	emptySpace = color.New(color.FgWhite)
	red        = color.New(color.FgRed)
	yellow     = color.New(color.FgYellow)
)

type Game struct {
	n       int
	players []*color.Color
	board   [][]*color.Color
}

func New(rows, columns, n int) (*Game, error) {
	if rows < 1 || columns < 1 {
		return nil, fmt.Errorf("rows and columns must be greater than 0")
	}
	if rows < n || columns < n {
		return nil, fmt.Errorf("rows and columns must be greater than or equal to n")
	}
	if n <= 0 {
		return nil, fmt.Errorf("n must be greater than zero")
	}

	b := make([][]*color.Color, rows)
	for i := range b {
		b[i] = make([]*color.Color, columns)
		for j := range b[i] {
			b[i][j] = emptySpace
		}
	}

	g := &Game{
		players: []*color.Color{red, yellow},
		board:   b,
		n:       n,
	}

	return g, nil
}

func (g *Game) Play() {
	stdin := bufio.NewReader(os.Stdin)
	currentPlayer := 0
	turns := 0

	g.DisplayBoard()

	for {
		var row, col int
		// Prompt for column to play
		for {
			fmt.Printf("Choose a column to play: ")
			if _, err := fmt.Scanf("%d", &col); err != nil {
				fmt.Printf("Invalid input: %v\n", err)
				stdin.ReadString('\n')
				continue
			}

			r, err := g.placeMarker(g.players[currentPlayer], col)
			if err != nil {
				fmt.Printf("Invalid move: %v\n", err)
				continue
			}
			row = r
			fmt.Println()
			break
		}

		g.DisplayBoard()

		if win, err := g.checkWin(g.players[currentPlayer], row, col); err != nil {
			log.Fatalf("checkWin failed: %v\n", err)
		} else if win {
			fmt.Printf("Player %d won!\n", currentPlayer+1)
			break
		}

		turns++
		if turns == len(g.board)*len(g.board[0]) {
			fmt.Println("Game ended in draw")
			break
		}

		currentPlayer = (currentPlayer + 1) % len(g.players)
	}
}

// placeMarker returns row index where marker was placed
func (g *Game) placeMarker(c *color.Color, col int) (int, error) {
	if c == nil {
		return 0, fmt.Errorf("c must not be nil")
	}
	if numCols := len(g.board[0]); col < 0 || col >= numCols {
		return 0, fmt.Errorf("col must be between 0 and %d", numCols-1)
	}

	for i := len(g.board) - 1; i >= 0; i-- {
		if g.board[i][col] == emptySpace {
			g.board[i][col] = c
			return i, nil
		}
	}
	return 0, fmt.Errorf("column is full")
}

func (g Game) checkWin(c *color.Color, row, col int) (bool, error) {
	counted := 1

	// Check for continuous horizontal line of g.n or more of color
	for i := col - 1; i >= 0; i-- {
		if g.board[row][i] != c {
			break
		}
		counted++
	}

	for i := col + 1; i < len(g.board[0]); i++ {
		if g.board[row][i] != c {
			break
		}
		counted++
	}

	if counted >= g.n {
		return true, nil
	}
	counted = 1

	// Check for continuous vertical line of g.n or more of color
	for i := row - 1; i >= 0; i-- {
		if g.board[i][col] != c {
			break
		}
		counted++
	}

	for i := row + 1; i < len(g.board); i++ {
		if g.board[i][col] != c {
			break
		}
		counted++
	}

	if counted >= g.n {
		return true, nil
	}
	counted = 1

	// Check for continuous incline of g.n or more of color
	i := row + 1
	j := col - 1
	for {
		if i >= len(g.board) || j < 0 || g.board[i][j] != c {
			break
		}
		counted++
		i++
		j--
	}

	i = row - 1
	j = col + 1
	for {
		if i < 0 || j >= len(g.board[0]) || g.board[i][j] != c {
			break
		}
		counted++
		i--
		j++
	}

	if counted >= g.n {
		return true, nil
	}
	counted = 1

	// Check for continuous decline of g.n or more of color
	i = row - 1
	j = col - 1
	for {
		if i < 0 || j < 0 || g.board[i][j] != c {
			break
		}
		counted++
		i--
		j--
	}

	i = row + 1
	j = col + 1
	for {
		if i >= len(g.board) || j >= len(g.board[0]) || g.board[i][j] != c {
			break
		}
		counted++
		i++
		j++
	}
	if counted >= g.n {
		return true, nil
	}
	return false, nil
}

func (g Game) DisplayBoard() {
	for _, row := range g.board {
		fmt.Printf(" ")
		for _, c := range row {
			c.Printf(marker + " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
