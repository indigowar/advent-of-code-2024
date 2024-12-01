package first

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func readFile(filename string) ([]int, []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	left := make([]int, 0)
	right := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		var l, r int
		fmt.Sscanf(line, "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}

	return left, right
}

func One() {
	left, right := readFile("./inputs/first.txt")

	slices.Sort(left)
	slices.Sort(right)

	distance := 0
	for i := range left {
		distance += max(left[i], right[i]) - min(left[i], right[i])
	}

	fmt.Println("1.1 Distance: ", distance)
}
