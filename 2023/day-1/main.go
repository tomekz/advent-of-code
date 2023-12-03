// https://adventofcode.com/2023/day/1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var digitWords = map[string]byte{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

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
			} else {
				word := ""
				for j := i; j < len(line); j++ {
					word += string(line[j])
					if v, ok := digitWords[word]; ok {
						digits = append(digits, v)
						i += len(word) - 1
						break
					}
				}
			}
		}
		first := digits[0]
		last := digits[len(digits)-1]
		sumStr := string(first) + string(last)

		sum, _ := strconv.Atoi(sumStr)

		// print the sum of the first and last digits
		fmt.Printf("%s %s %s %s\n", string(digits), string(first), string(last), sumStr)
		total += sum
	}

	fmt.Println(total)
}
