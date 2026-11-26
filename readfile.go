package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func readFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func splitBlocks(content string) ([]Tetromino, error) {
	var blocks [][]string
	var current []string
	var count int
	var cells []Point
	var tetromino Tetromino
	var pieces []Tetromino

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 4 && strings.TrimSpace(line) != "" {
			return nil, errors.New("ERROR")
		}

		if strings.TrimSpace(line) == "" {
			if len(blocks) < 27 && len(current) == 4 && count == 4 {
				blocks = append(blocks, current)
				current = nil
				count = 0
				continue
			} else {
				return nil, errors.New("ERROR")
			}
		}
		for _, ch := range line {
			if ch == '#' {
				count++
			} else if ch == '.' {
				continue
			} else {
				return nil, errors.New("ERROR")
			}

		}
		current = append(current, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(current) > 0 {
		blocks = append(blocks, current)
	}
	for _, block := range blocks {
		for r, line := range block {
			for c, ch := range line {
				if ch == '#' {
					cells = append(cells, Point{Row: r, Col: c})
				}
			}
		}
		tetromino = Tetromino{Cells: cells}
		pieces = append(pieces, tetromino)
		for i, p := range pieces {
			pieces[i] = Tetromino{
				Cells:  p.Cells,
				Letter: 'A' + rune(i),
			}
		}
		cells = nil
	}
	return pieces, nil
}
