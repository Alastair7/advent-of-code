package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}

func getGameID(line string) int {
	colonIndex := strings.Index(line, ":")
	whiteSpaceIndex := strings.Index(line, " ")

	gameID, _ := strconv.Atoi((line[whiteSpaceIndex + 1 : colonIndex]))
	return gameID
}

func areSetsValid(scoreMap map[string]int64, gameSets []string) bool {

	var isValid bool = true

	for _, gameSet := range gameSets {
		scoreMap["red"] = 0 
		scoreMap["blue"] = 0
		scoreMap["green"] = 0

		points := strings.Split(gameSet, ",")

		for _, point := range points {
			if strings.Contains(point, "red") {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				scoreMap["red"] += digit
			} else if strings.Contains(point, "blue") {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				scoreMap["blue"] += digit
			} else {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				scoreMap["green"] += digit
			}
		}

		if (scoreMap["red"] > int64(12)) || (scoreMap["blue"] > int64(14) || scoreMap["green"] > int64(13) ) {
			isValid = false
		}
	}

	return isValid
}
func isGamePossible(line string) bool {
	var isPossible bool = true

	scoreMap := make(map[string]int64)

	colonIndex := strings.Index(line, ":")
	formattedLine := strings.Trim(line[colonIndex + 1:], " ")

	gameSets := strings.Split(formattedLine, ";")

	isPossible = areSetsValid(scoreMap, gameSets)

	return isPossible
}

func getGameFewestOptions(line string) [3] int64 {
	colonIndex := strings.Index(line, ":")
	formattedLine := strings.Trim(line[colonIndex + 1:], " ")

	fewestOptions := [3] int64 {0, 0 ,0}

	gameSets := strings.Split(formattedLine, ";")

	for _, gameSet := range gameSets {
		points := strings.Split(gameSet, ",")

		for _, point := range points {
			if strings.Contains(point, "red") {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				if digit > fewestOptions[0] {
					fewestOptions[0] = digit
				}
			} else if strings.Contains(point, "blue") {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				if digit > fewestOptions[1] {
					fewestOptions[1] = digit
				}
			} else {
				formatPoint := strings.Trim(point, " ")

				whiteSpaceIndex := strings.Index(formatPoint, " ")
				digit, _ := strconv.ParseInt(formatPoint[0 : whiteSpaceIndex], 10, 64)

				if digit > fewestOptions[2] {
					fewestOptions[2] = digit
				}
			}
		}
	}

	return fewestOptions
}

func part1(file *os.File) {
	scanner := bufio.NewScanner(file)
	var result int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		gameID := getGameID(line)

		isPossible := isGamePossible(line)

		if isPossible {
			result += int64(gameID)
		}

	}

	fmt.Println(result)

	if err := scanner.Err(); err != nil {
		log.Fatal((err))
	}
}

func part2(file *os.File) {
	scanner := bufio.NewScanner(file)
	var result int64 = 0
	fewestOptions := make(map[int] [3]int64)

	for scanner.Scan() {
		line := scanner.Text()

		gameID := getGameID(line)

		gameMaxPoints := getGameFewestOptions(line)

		fewestOptions[gameID] = gameMaxPoints
	}

	for _,gamePoints := range fewestOptions {
		result += gamePoints[0] * gamePoints[1] * gamePoints[2] 
	} 

	fmt.Println(result)

}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	

	// Part 1 FUNC
	//part1(file)

	// Part 2 FUNC
	//part2(file)
}