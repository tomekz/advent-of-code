// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func parseInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	input := make([]string, 0)
	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		input = append(input, string(line))
	}

	return input, nil
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func main() {
	input, err := parseInput("input")
	if err != nil {
		fmt.Println(err)
		return
	}

	var total int
	for _, line := range input {
		digits := make([]byte, 0)
		for i := 0; i < len(line); i++ {
			if isDigit(line[i]) {
				digits = append(digits, line[i])
			}
		}
		first := digits[0]
		last := digits[len(digits)-1]
		sumStr := string(first) + string(last)

		sum, _ := strconv.Atoi(sumStr)

		total += sum
	}

	fmt.Println(total)
}
