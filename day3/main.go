package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x int
	y int
}

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
	gear := make(map[point][]int)
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
			part := false
			if (j-1 >= 0) && !(arr[i][j-1] >= '0' && arr[i][j-1] <= '9') && !(arr[i][j-1] == '.') {
				ans += num
				part = true
			}
			if (end < len(value)) && !(arr[i][end] >= '0' && arr[i][end] <= '9') && !(arr[i][end] == '.') {
				ans += num
				part = true
			}
			if i-1 >= 0 {
				k := j - 1
				for ; k <= end; k++ {
					if k < 0 || k >= len(value) {
						continue
					}
					if (arr[i-1][k] != '.') && !(arr[i-1][k] >= '0' && arr[i-1][k] <= '9') {
						ans += num
						part = true
						break
					}
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
						part = true
						break
					}
				}

			}
			if part {
				if (j-1 >= 0) && arr[i][j-1] == '*' {
					gear[point{x: i, y: j - 1}] = append(gear[point{x: i, y: j - 1}], num)
				}
				if (end < len(value)) && arr[i][end] == '*' {
					gear[point{x: i, y: end}] = append(gear[point{x: i, y: end}], num)
				}
				if i-1 >= 0 {
					k := j - 1
					for ; k <= end; k++ {
						if k < 0 || k >= len(value) {
							continue
						}
						if arr[i-1][k] == '*' {
							gear[point{x: i - 1, y: k}] = append(gear[point{x: i - 1, y: k}], num)
						}
					}

				}
				if i+1 < len(arr) {
					k := j - 1
					for ; k <= end; k++ {
						if k < 0 || k >= len(value) {
							continue
						}
						if arr[i+1][k] == '*' {
							gear[point{x: i + 1, y: k}] = append(gear[point{x: i + 1, y: k}], num)
						}
					}

				}
				j = end
				continue
			}
			j++

		}
	}
	power := 0
	for _, value := range gear {
		if len(value) == 2 {
			power += (value[0] * value[1])
		}
	}

	fmt.Println(power)
}
