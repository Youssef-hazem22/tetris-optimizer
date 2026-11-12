package main

func normalizeAll(pieces []Tetromino) []Tetromino {
	result := make([]Tetromino, 0, len(pieces))
	for _, p := range pieces {
		result = append(result, normalize(p))
	}
	return result
}

func normalize(t Tetromino) Tetromino {
	minRow, minCol := t.Cells[0].Row, t.Cells[0].Col
	for _, p := range t.Cells {
		if p.Row < minRow {
			minRow = p.Row
		}
		if p.Col < minCol {
			minCol = p.Col
		}
	}

	cells := make([]Point, len(t.Cells))
	for i, p := range t.Cells {
		cells[i] = Point{Row: p.Row - minRow, Col: p.Col - minCol}
	}
	return Tetromino{Cells: cells, Letter: t.Letter}
}
