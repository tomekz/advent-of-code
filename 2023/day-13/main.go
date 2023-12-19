package main

import (
	"bufio"
	"os"
)

type Pattern [][]byte

func parseFile(file *os.File) []Pattern {
	var patterns []Pattern
	scanner := bufio.NewScanner(file)

	tmp := make([][]byte, 0)
	for scanner.Scan() {
		line := scanner.Text()
		chars := []byte(line)
		if line == "" {
			patterns = append(patterns, tmp)
			tmp = make([][]byte, 0)
			continue
		} else {
			tmp = append(tmp, chars)
		}
	}
	patterns = append(patterns, tmp)

	return patterns
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	patterns := parseFile(file)
	for _, pattern := range patterns {
		for _, line := range pattern {
			for _, char := range line {
				print(string(char))
			}
			println()
		}
		println()
	}
}
