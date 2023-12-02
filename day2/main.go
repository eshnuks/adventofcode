package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getSubGameCount(subGame string) (int, int, int) {
	// fmt.Println(subGame)
	red := 0
	blue := 0
	green := 0
	arr := strings.Split(subGame, ", ")
	for _, value := range arr {
		parts := strings.Split(value, " ")
		count, _ := strconv.Atoi(parts[0])
		if parts[1] == "red" {
			red = count
		} else if parts[1] == "blue" {
			blue = count
		} else {
			green = count
		}

	}

	return red, blue, green
}

func getGamePower(game string) int {
	// fmt.Println(game)
	red := 0
	blue := 0
	green := 0
	arr := strings.Split(game, "; ")
	for _, value := range arr {
		redTemp, blueTemp, greenTemp := getSubGameCount(value)
		red = max(red, redTemp)
		blue = max(blue, blueTemp)
		green = max(green, greenTemp)
	}
	return red * blue * green
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
		parts := strings.Split(line, ": ")
		ans += getGamePower(parts[1])
	}

	fmt.Println(ans)
}
