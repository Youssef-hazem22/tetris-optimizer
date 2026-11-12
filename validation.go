package main

import (
	"errors"
	"fmt"
)

const maxPieces = 26

func validateBlocks(blocks [][]string) error {
	if len(blocks) == 0 {
		return errors.New("no tetrominoes found")
	}
	if len(blocks) > maxPieces {
		return fmt.Errorf("too many tetrominoes: %d (max %d)", len(blocks), maxPieces)
	}

	for i, block := range blocks {
		if err := validateBlock(block); err != nil {
			return fmt.Errorf("block %d: %w", i+1, err)
		}
	}
	return nil
}

func validateBlock(block []string) error {
	if len(block) != 4 {
		return fmt.Errorf("expected 4 lines, got %d", len(block))
	}

	for r, line := range block {
		if len(line) != 4 {
			return fmt.Errorf("line %d: expected 4 characters, got %d", r+1, len(line))
		}
		for c, ch := range line {
			if ch != '#' && ch != '.' {
				return fmt.Errorf("line %d, col %d: invalid character %q", r+1, c+1, ch)
			}
		}
	}
	return nil
}

func validateShapes(pieces []Tetromino) error {
	for i, p := range pieces {
		if err := validateShape(p); err != nil {
			return fmt.Errorf("piece %d: %w", i+1, err)
		}
	}
	return nil
}

func validateShape(t Tetromino) error {
	if len(t.Cells) != 4 {
		return fmt.Errorf("expected 4 '#', got %d", len(t.Cells))
	}

	set := make(map[Point]bool, len(t.Cells))
	for _, p := range t.Cells {
		set[p] = true
	}

	directions := []Point{{Row: -1, Col: 0}, {Row: 1, Col: 0}, {Row: 0, Col: -1}, {Row: 0, Col: 1}}
	adjacencies := 0
	for _, p := range t.Cells {
		for _, d := range directions {
			neighbor := Point{Row: p.Row + d.Row, Col: p.Col + d.Col}
			if set[neighbor] {
				adjacencies++
			}
		}
	}

	if adjacencies != 6 && adjacencies != 8 {
		return fmt.Errorf("cells are not connected (adjacencies = %d)", adjacencies)
	}
	return nil
}
