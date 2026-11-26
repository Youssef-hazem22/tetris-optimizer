package main

import (
	"fmt"
	"os"
)

type Point struct {
	Row, Col int
}

type Tetromino struct {
	Cells  []Point
	Letter rune
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(" len ERROR ")
		os.Exit(1)
	}

	content, err := readFile(os.Args[1])
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces, err := splitBlocks(content)
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	if pieces, err = Allmatching(pieces); err != nil {
		fmt.Println("ERROR")
		return
	}
	n := minBoardSize(pieces)
	for {
		board := newBoard(n)
		used := make([]bool, len(pieces))
		waste := n*n - 4*len(pieces)
		if solve(board, pieces, used, 0, len(pieces), waste) {
			fmt.Print(boardString(board))
			return
		}
		n++
	}

}
