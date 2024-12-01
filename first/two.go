package first

import (
	"fmt"
)

func Two() {
	left, right := readFile("./inputs/first.txt")

	rightOccurrences := make(map[int]int, 0)
	for _, v := range right {
		c, _ := rightOccurrences[v]
		rightOccurrences[v] = c + 1
	}

	score := 0
	for _, v := range left {
		c, _ := rightOccurrences[v]
		score += c * v
	}

	fmt.Println("1.2 Score: ", score)
}
