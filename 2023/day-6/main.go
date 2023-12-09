package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time     int
	distance int
}

func (r *Race) calculateWaysToWin() int {
	ways := 0
	for i := 0; i < r.time; i++ {
		maxDistance := i * (r.time - i)
		if maxDistance > r.distance {
			ways++
		}
	}
	return ways
}

func parseInput() []Race {
	result := []Race{}
	content, _ := os.ReadFile("input")
	lines := strings.Split(string(content), "\n")

	numberMatcher := regexp.MustCompile(`\d+`)
	times := numberMatcher.FindAllString(lines[0], -1)
	distances := numberMatcher.FindAllString(lines[1], -1)

	for i := 0; i < len(times); i++ {
		time, _ := strconv.Atoi(times[i])
		distance, _ := strconv.Atoi(distances[i])
		result = append(result, Race{time, distance})
	}

	return result
}

func main() {
	races := parseInput()

	megaRace := Race{}
	megaDistance := ""
	megaTime := ""
	totalWays := 1
	for _, race := range races {
		totalWays *= race.calculateWaysToWin()
		megaDistance += strconv.Itoa(race.distance)
		megaTime += strconv.Itoa(race.time)
		megaRace.distance, _ = strconv.Atoi(megaDistance)
		megaRace.time, _ = strconv.Atoi(megaTime)
	}

	fmt.Println(totalWays)
	fmt.Println(megaRace.calculateWaysToWin())
}
