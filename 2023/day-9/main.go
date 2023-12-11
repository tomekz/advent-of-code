package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func subtractItemsInSlice(slice []int) []int {
	var result []int
	for i := 0; i < len(slice)-1; i++ {
		result = append(result, slice[i+1]-slice[i])
	}
	return result
}

func predictNextItem(slice []int) int {
	allZero := false
	result := make([][]int, 0)
	result = append(result, slice)
	currSlice := slice
	for {
		if allZero {
			break
		}
		currSlice = subtractItemsInSlice(currSlice)
		result = append(result, currSlice)

		total := 0
		for _, v := range currSlice {
			total += v
		}
		if total == 0 {
			allZero = true
		}
	}

	nextValue := 0
	for i := 0; i < len(result); i++ {
		lastItem := result[i][len(result[i])-1]
		nextValue += lastItem
	}
	return nextValue
}

func main() {
	file, _ := os.Open("input")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")

		history := make([]int, 0)
		for _, v := range text {
			num, _ := strconv.Atoi(v)
			history = append(history, num)
		}

		nextValue := predictNextItem(history)
		fmt.Println(nextValue)
		total += nextValue
	}
	fmt.Printf("Total: %d\n", total)
}
