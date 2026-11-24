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

	blocks, err := splitBlocks(content)
	if err != nil {
		fmt.Println("ERROR")
		return
	}

	if err = validateBlocks(blocks); err != nil {
		fmt.Println("ERROR")
		return
	}

	pieces := parseTetrominoes(blocks)

	if pieces, err = validateShapes(pieces); err != nil {
		fmt.Println("ERROR")
		return
	}
	pieces = assignLetters(pieces)
	n := minBoardSize(pieces)
	for {
		board := newBoard(n)
		if solve(board, pieces, 0) {
			fmt.Print(boardString(board))
			return
		}
		n++
	}

}
