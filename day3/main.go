package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	ans := 0
	var arr []string
	// Iterate through each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, line)
	}
	for i, value := range arr {
		for j := 0; j < len(value); {
			if !(value[j] >= '0' && value[j] <= '9') {
				j++
				continue
			}
			num := 0
			end := j
			for ; end < len(value) && value[end] >= '0' && value[end] <= '9'; end++ {
				num = num*10 + int(value[end]-'0')
			}
			if (j-1 >= 0) && !(arr[i][j-1] >= '0' && arr[i][j-1] <= '9') && !(arr[i][j-1] == '.') {
				ans += num
				j = end
				continue
			}
			if (end < len(value)) && !(arr[i][end] >= '0' && arr[i][end] <= '9') && !(arr[i][end] == '.') {
				ans += num
				j = end
				continue
			}
			if i-1 >= 0 {
				k := j - 1
				for ; k <= end; k++ {
					if k < 0 || k >= len(value) {
						continue
					}
					if (arr[i-1][k] != '.') && !(arr[i-1][k] >= '0' && arr[i-1][k] <= '9') {
						ans += num
						break
					}
				}
				if k <= end {
					j = end
					continue
				}

			}
			if i+1 < len(arr) {
				k := j - 1
				for ; k <= end; k++ {
					if k < 0 || k >= len(value) {
						continue
					}
					if (arr[i+1][k] != '.') && !(arr[i+1][k] >= '0' && arr[i+1][k] <= '9') {
						ans += num
						break
					}
				}
				if k <= end {
					j = end
					continue
				}

			}
			j++

		}
	}

	fmt.Println(ans)
}
