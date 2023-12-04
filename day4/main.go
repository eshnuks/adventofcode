package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	ans := 0
	cardNum := 1
	// Iterate through each line in the file
	cards := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.FieldsFunc(line, unicode.IsSpace)
		start := 2
		cards[cardNum]++
		winNum := make(map[string]bool)
		for ; arr[start] != "|"; start++ {
			winNum[arr[start]] = true
		}
		start++
		score := 0
		for ; start < len(arr); start++ {
			if winNum[arr[start]] {
				score++
			}
		}
		// fmt.Println(score)
		for k := cardNum + 1; k <= cardNum+score; k++ {
			cards[k] += (cards[cardNum])
		}
		cardNum++

	}
	for _, value := range cards {
		ans += value
	}

	fmt.Println(ans)
}
