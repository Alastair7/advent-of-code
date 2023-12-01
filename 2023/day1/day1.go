package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}


func getNumbers(line string, numberMap map[string]string) []int {
	digits := []int{}
	
	for word,value := range numberMap {
		if strings.Contains(line, word) {
			line = strings.ReplaceAll(line, word, value)
		}
	}
	
	for _, char := range line {
		if unicode.IsDigit(char) {
			digit := int(char - '0')
			digits = append(digits, digit)
		}
	}
	return digits
}

// PART 1 FUNCTION
/*
func handleDigits(line string) []int{
	digits := []int{}

	for _, char := range line {
		if unicode.IsDigit(char) {
			digit := int(char - '0')
			digits = append(digits, digit)
		}
	}

	return digits
}
*/

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	var result int64 = 0
	
	numberMap := make(map[string]string)

	numberMap["one"] = "o1e"
	numberMap["two"] = "t2o"
	numberMap["three"] = "t3e"
	numberMap["four"] = "f4r"
	numberMap["five"] = "f5e"
	numberMap["six"] = "s6x"
	numberMap["seven"] = "s7n"
	numberMap["eight"] = "e8t"
	numberMap["nine"] = "n9e"



	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		digits := getNumbers(line, numberMap)

		unionDigitsString := strconv.Itoa(digits[0]) + strconv.Itoa(digits[len(digits) - 1])
		unionDigits, _ := strconv.ParseInt(unionDigitsString, 10, 64)

		result += unionDigits
	}

	if err := scanner.Err(); err != nil {
		log.Fatal((err))
	}

	fmt.Println(result)
}