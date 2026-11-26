go build -o tetris-optimizer .
for f in testcases/g*.txt testcases/b*.txt testcases/hard.txt; do echo "== $f"; ./tetris-optimizer "$f"; done