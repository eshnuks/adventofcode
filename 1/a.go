package main

import (
	"bufio"
	"fmt"
	"os"
)

func checkSubstring(line string) int {
	if len(line) >= 4 && line[0:4] == "zero" {
		return 0
	}
	if len(line) >= 3 && line[0:3] == "one" {
		return 1
	}
	if len(line) >= 3 && line[0:3] == "two" {
		return 2
	}
	if len(line) >= 5 && line[0:5] == "three" {
		return 3
	}
	if len(line) >= 4 && line[0:4] == "four" {
		return 4
	}
	if len(line) >= 4 && line[0:4] == "five" {
		return 5
	}
	if len(line) >= 3 && line[0:3] == "six" {
		return 6
	}
	if len(line) >= 5 && line[0:5] == "seven" {
		return 7
	}
	if len(line) >= 5 && line[0:5] == "eight" {
		return 8
	}
	if len(line) >= 4 && line[0:4] == "nine" {
		return 9
	}
	return -1
}

func calibrateText(line string) int {
	first := -1
	second := -1
	for i, value := range line {
		letterDigit := checkSubstring(line[i:])
		if value >= '0' && value <= '9' {
			if first == -1 {
				first = (int)(value - '0')
			}
			second = (int)(value - '0')
		} else if letterDigit != -1 {
			if first == -1 {
				first = letterDigit
			}
			second = letterDigit
		}
	}

	if first == -1 {
		return 0
	}
	return 10*first + second

}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	ans := 0

	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		ans += calibrateText(line)
	}

	fmt.Println(ans)

}
