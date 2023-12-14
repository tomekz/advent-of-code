package main

import (
	"bufio"
	"fmt"
	"os"
)

type Galaxy struct {
	x int
	y int
}

func parseInput() []Galaxy {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	result := make([]Galaxy, 0)
	emptyColumns := make([]bool, 0)
	y := 0
	distance := 1

	for scanner.Scan() {
		text := scanner.Text()
		row := make([]string, len(text))
		empty := true
		for x, char := range text {
			if y == 0 {
				emptyColumns = append(emptyColumns, true)
			}
			if string(char) == "#" {
				galaxy := Galaxy{x, y}
				result = append(result, galaxy)
				empty = false
				emptyColumns[x] = false
			}
			row[x] = string(char)
		}
		if empty {
			y += distance
		}
		y++

	}

	for i := range result {
		n := result[i].x
		for j := 0; j < result[i].x; j++ {
			if emptyColumns[j] {
				n += distance
			}
		}
		result[i].x = n
	}

	return result
}

func CalcDistances(nodeMap []Galaxy) int {
	dist := 0
	for i := range nodeMap {
		for j := i + 1; j < len(nodeMap); j++ {
			dist += Abs(nodeMap[i].x-nodeMap[j].x) + Abs(nodeMap[i].y-nodeMap[j].y)
		}
	}

	return dist
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	galaxy := parseInput()
	fmt.Println(galaxy)
	total := CalcDistances(galaxy)
	fmt.Println(total)
}
