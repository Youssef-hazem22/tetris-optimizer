# Tetris Optimizer

A small program that packs Tetris pieces into the **smallest square box possible**.

Written in Go, using only the language's built-in tools (no outside libraries).

## What problem does it solve?

Think of packing a suitcase with oddly shaped items. You want the smallest
suitcase that still fits everything, and you don't want to waste space.

This program does the same with Tetris pieces. You give it a list of pieces,
and it finds the smallest square that holds them all. Each piece is marked with
its own letter (the first piece is `A`, the second is `B`, and so on), so you
can see where everything ended up. A dot (`.`) is an empty spot.

## See it in action

**Example 1.** A straight line (`A`) and a square (`B`):

```
ABB.
ABB.
A...
A...
```

**Example 2.** Eight different pieces, packed into a 6 x 6 square:

```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

Eight pieces cover 32 squares, and the box has 36. So only 4 spots are empty,
and no smaller box could work.

## Why it is harder than it looks

With a handful of pieces, you can try every arrangement. But the number of
possible arrangements grows explosively with each new piece. My first version
worked well for small sets and then **froze at around 12 pieces**.

So I measured where the time was going and redesigned the search to skip
arrangements that are clearly hopeless, long before trying them. After that,
most sets of 20 to 26 pieces are solved in **a few seconds**.

## It does not crash on bad input

If the file is wrong in any way, the program prints one word, `ERROR`, instead
of crashing or giving a strange answer. For example:

- the file is missing or empty
- a piece has the wrong size, or contains a wrong character
- a piece is broken into two separate parts
- pieces are not separated by an empty line
- there are too many pieces

I tested this with 18 different kinds of broken files and with automatically
generated random input. None of them make the program crash.

## What this project shows

| Skill | How it shows up here |
|---|---|
| Problem solving | Turning a puzzle into a step-by-step method a computer can follow |
| Algorithms | A search that tries options, backs up when stuck, and skips hopeless paths |
| Performance | Finding the slow part by measuring, then redesigning it (from freezing at 12 pieces to seconds at 20+) |
| Testing | Automated tests that cover about 96% of the code, plus random-input testing |
| Reliability | Every kind of bad input is handled, with a clear result |
| Clean code | Small single-purpose functions, one place that handles all errors |

## Try it yourself

You need [Go](https://go.dev/dl/) installed. Then, in this folder:

```bash
go run . sample.txt     # solve the example file
go test ./...           # run the automated tests
```

`sample.txt` is the example with eight pieces. A piece is drawn with `#`
(filled) and `.` (empty) in a 4 x 4 grid, and pieces are separated by an empty
line:

```
##..
.##.
....
....
```

## Good to know

- Pieces are placed as they are drawn; they are never rotated.
- The program accepts up to 26 pieces (one letter each).
- Finding the very smallest box is a genuinely hard problem. Most 20 to 26 piece
  sets finish in a few seconds, but some unlucky ones take much longer.
  (In my tests, 9 of 10 random sets finished in under 4 seconds, and one took
  over 12.)

---

## For developers

- **Language:** Go 1.18 or newer, standard library only.
- **Files:** `main.go` (program), `main_test.go` (tests), `stress.sh` (runs the
  real program on 18 broken inputs), `sample.txt` (example).
- **Approach:** validate the input, move every piece to the top-left corner,
  then search for the smallest square. A fast cell-by-cell search (`packer`)
  finds the smallest size, and a piece-by-piece backtracking search (`solver`)
  builds the final board. Both skip hopeless branches (an empty area of `s`
  cells wastes at least `s mod 4` cells; identical pieces are not tried in
  every order).
- **Errors:** every failure is a named error (`ErrBadFormat`, `ErrBadShape`, ...)
  handled in one place. `main()` only prints `ERROR` and exits with status 1.
  Set `TETRIS_DEBUG=1` to see the reason on standard error.
- **Tests:** table-driven unit tests, a fuzz test, benchmarks and a CPU profile
  workflow. See `go test -cover ./...`.
