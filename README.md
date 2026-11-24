# Tetris Optimizer

A small program that packs Tetris pieces into the **smallest square box possible**.

Written in Go, using only the language's built-in tools (no outside libraries).

## What problem does it solve?

Think of packing a suitcase with oddly shaped items. You want the smallest
suitcase that still fits everything, with as little wasted space as possible.

This program does the same with Tetris pieces. You give it a file with a list
of pieces, and it finds the smallest square that holds them all. Every piece is
marked with its own letter (the first piece is `A`, the second is `B`, and so
on), so you can see where each one ended up. A dot (`.`) is an empty spot.

## See it in action

**Example 1.** Four pieces (a vertical line, a horizontal line, an L-shaped
piece and a zig-zag piece) fit in a 5 x 5 square:

```
ABBBB
ACCC.
A..C.
ADD..
DD...
```

**Example 2.** Eight pieces (`sample.txt`) fit in a 6 x 6 square:

```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

Eight pieces cover 32 squares and the box has 36, so only 4 spots stay empty.
A smaller box could not work, because it would have fewer than 32 squares.

## How to use it

You need [Go](https://go.dev/dl/) installed. Then, in this folder:

```bash
go run . sample.txt
```

The program takes **exactly one argument**: the path to a text file with the
pieces.

### How a piece is written

Each piece is drawn in a 4 x 4 grid: `#` is a filled square and `.` is an empty
one. Pieces are separated by one empty line.

```
...#
...#
...#
...#

....
.##.
.##.
....
```

The `#` can be anywhere inside the grid. The program moves each piece to the
corner by itself. Files with Windows line endings work too.

## If the file is wrong, it prints `ERROR`

Instead of crashing or giving a strange answer, the program prints `ERROR` when:

- the file cannot be opened, or it is empty
- a piece is not exactly 4 lines of 4 characters
- a piece contains anything other than `#` and `.`
- a piece does not have exactly 4 `#`
- the 4 `#` of a piece are not all connected (touching diagonally does not count)
- pieces are not separated by one empty line (for example, two empty lines in a row)
- there are more than 26 pieces (there are only 26 letters)

## Example files included

The folder has ready-made input files to try:

| File | What it is | Result |
|---|---|---|
| `sample.txt` | 8 pieces (the example from the project brief) | 6 x 6 square |
| `g0.txt` | 1 piece (a square) | 2 x 2 square |
| `g1.txt` | 4 pieces | 5 x 5 square |
| `g2.txt` | 8 pieces | 6 x 6 square |
| `g3.txt` | 11 pieces | 7 x 7 square |
| `b0.txt` | a piece with 5 filled squares | `ERROR` |
| `b1.txt` | a piece whose squares only touch diagonally | `ERROR` |
| `b2.txt` | a piece split into two separate parts | `ERROR` |
| `b3.txt` | a piece with no filled squares at all | `ERROR` |
| `b4.txt` | another piece split into two parts | `ERROR` |
| `b5.txt` | pieces not separated correctly | `ERROR` |
| `hard.txt` | 12 pieces, a tough one | 7 x 7 square, but slow (see below) |

(`g` is for good input, `b` is for bad input.)

To run all of them at once (Linux, macOS or Git Bash):

```bash
go build -o tetris-optimizer .
for f in g*.txt b*.txt; do echo "== $f"; ./tetris-optimizer "$f"; done
```

## What this project shows

| Skill | How it shows up here |
|---|---|
| Problem solving | Turning a puzzle into a step-by-step method a computer can follow |
| Algorithms | A search that tries an option, backs up when it gets stuck, and tries the next one (backtracking) |
| Careful input checking | Every kind of broken file ends in a clear `ERROR` |
| Organized code | Each file has one job: reading, checking, converting, building the board, solving |
| Testing by example | A set of good, bad and hard input files |

## Good to know

- Pieces are placed exactly as drawn. They are never rotated.
- Up to 26 pieces are accepted.
- Speed depends on how many pieces there are. The included sets of up to 11
  pieces finish almost instantly. `hard.txt` (12 pieces) took about 25 seconds
  in a test run, and bigger sets can take much longer. The search tries every
  option without shortcuts, and the number of options grows very quickly with
  each new piece.

## Next steps

- Add automated tests (`go test`) for the checking and solving steps.
- Make the search faster by skipping arrangements that clearly cannot work.

---

## For developers

- **Language:** Go, standard library only. `go.mod` declares Go 1.26.2.
- **Smallest square:** the board starts small and grows by one row and one
  column every time the pieces do not fit. The first size where everything fits
  is therefore the smallest one.

### Files

| File | Job |
|---|---|
| `main.go` | Entry point; runs the steps in order and prints `ERROR` on any failure |
| `readfile.go` | Reads the file and splits it into blocks (a line-by-line scanner, so `\r\n` works) |
| `validation.go` | Checks block layout, checks each shape, moves each piece to the top-left corner |
| `parsing.go` | Turns each block into the coordinates of its `#`, and gives each piece its letter |
| `buildBoard.go` | Starting board size, creating the board, turning it into text |
| `solver.go` | `canPlace`, `place`, `remove` and the recursive `solve` (backtracking) |

### How it works

```
read file -> split into blocks -> check blocks -> parse '#' cells -> check shapes
          -> move to corner -> assign letters -> try board sizes with solve() -> print
```

- **Shape check:** a piece must have exactly 4 cells. For every cell, the code
  counts the filled cells directly above, below, left and right of it. A
  connected piece gives a total of 6 (three touching pairs, each counted twice)
  or 8 (the 2 x 2 square). Any other total means the cells are not all connected.
- **Solving:** `solve(board, pieces, i)` places piece `i` in the first spot (row
  by row) where it fits, then calls itself for piece `i+1`. If a later piece
  cannot be placed, it removes piece `i` and tries the next spot.
- **Board size:** if `solve` fails, `main` makes the board one cell bigger and
  starts again.