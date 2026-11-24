go build -o tetris-optimizer .
for f in g*.txt b*.txt hard.txt; do echo "== $f"; ./tetris-optimizer "$f"; done