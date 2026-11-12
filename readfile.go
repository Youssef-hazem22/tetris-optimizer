package main

import (
	"bufio"
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

func splitBlocks(content string) ([][]string, error) {
	var blocks [][]string
	var current []string

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		line := scanner.Text()

		if strings.TrimSpace(line) == "" {
			blocks = append(blocks, current)
			current = nil
			continue
		}
		current = append(current, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(current) > 0 {
		blocks = append(blocks, current)
	}
	return blocks, nil
}
