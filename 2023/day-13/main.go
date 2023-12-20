package main

import (
	"bufio"
	"fmt"
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

func colsMatch(m []string, c1, c2 int) bool {
	for i := 0; i < len(m); i++ {
		if m[i][c1] != m[i][c2] {
			return false
		}
	}
	return true
}

func reflectVertical(m []string, skip int) (int, bool) {
	width := len(m[0])
Outer:
	for i := 0; i < width-1; i++ {
		ti := i
		for j := i + 1; ; j++ {
			if !colsMatch(m, ti, j) {
				continue Outer
			}
			ti--
			if ti < 0 || j >= width-1 {
				if skip == i {
					continue Outer
				}
				return i, true
			}
		}
	}
	return 0, false
}

func reflectHorizontal(m []string, skip int) (int, bool) {
	height := len(m)
Outer:
	for i := 0; i < height-1; i++ {
		ti := i
		for j := i + 1; ; j++ {
			if m[ti] != m[j] {
				continue Outer
			}
			ti--
			if ti < 0 || j >= height-1 {
				if skip == i {
					continue Outer
				}
				return i, true
			}
		}
	}

	return 0, false
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

	s := 0
	for _, pattern := range patterns {
		lines := make([]string, len(pattern))
		for i, line := range pattern {
			lines[i] = string(line)
		}
		if v, ok := reflectVertical(lines, -1); ok {
			println("v", v)
			s += v + 1
		} else if h, ok := reflectHorizontal(lines, -1); ok {
			println("h", h)
			s += (v + 1) * 100
		} else {
			println("nope")
		}
	}
	fmt.Printf("There are %d valid patterns\n", s)
}
