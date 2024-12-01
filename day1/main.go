package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	return string(b)
}

func parseInput(i string) ([]int, []int) {
	splited := strings.Split(i, "\n")

	left := make([]int, len(splited))
	right := make([]int, len(splited))

	for _, s := range splited {
		nbs := strings.Split(s, "   ")

		leftValue, err := strconv.Atoi(nbs[0])
		if err != nil {
			panic(err)
		}

		rightValue, err := strconv.Atoi(nbs[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftValue)
		right = append(right, rightValue)
	}

	return left, right
}

func part1(left []int, right []int) {
	sort.Ints(left)
	sort.Ints(right)

	r := 0

	for i := 0; i < len(left); i++ {
		sum := left[i] - right[i]

		if sum < 0 {
			sum *= -1
		}

		r += sum
	}

	fmt.Printf("Part 1 : %d\n", r)
}

func part2(left []int, right []int) {
	m := make(map[int]int)

	r := 0

	for _, item := range right {
		if v, ok := m[item]; !ok {
			m[item] = 1
		} else {
			m[item] = v + 1
		}
	}

	for _, item := range left {
		r += m[item] * item
	}

	fmt.Printf("Part 2 : %d\n", r)
}

func main() {
	i := readFile()

	left, right := parseInput(i)

	part1(left, right)
	part2(left, right)
}
