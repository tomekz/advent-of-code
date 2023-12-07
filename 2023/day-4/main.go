package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseGameCard(card string) [2][]string {
	afterCard := strings.Split(card, ": ")[1]
	numberSplit := strings.Split(afterCard, "|")

	numberMatcher := regexp.MustCompile("[0-9]+")
	winningStrings := numberMatcher.FindAllString(numberSplit[0], -1)
	ourStrings := numberMatcher.FindAllString(numberSplit[1], -1)

	fmt.Println(afterCard)
	return [2][]string{winningStrings, ourStrings}
}

func GetNumbersFromString(numStr []string) ([]int, error) {
	nums := make([]int, len(numStr))
	for idx, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return make([]int, 0), err
		}
		nums[idx] = num
	}
	return nums, nil
}

func Part1(file io.Reader) (int, error) {
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		points := 0
		line := scanner.Text()

		gameStrings := ParseGameCard(line)

		winningNumbers, err := GetNumbersFromString(gameStrings[0])
		if err != nil {
			return -1, nil
		}
		winningMap := make(map[int]bool)
		for _, num := range winningNumbers {
			winningMap[num] = true
		}
		ourNumbers, err := GetNumbersFromString(gameStrings[1])
		if err != nil {
			return -1, nil
		}
		for _, val := range ourNumbers {
			_, ok := winningMap[val]
			if ok {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		totalPoints += points
	}
	return totalPoints, nil
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	points, err := Part1(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(points)
}
