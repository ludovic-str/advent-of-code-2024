package main

import (
	"fmt"
	"math"
	"os"
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

func parseInput(i string) [][]int {
	splited := strings.Split(i, "\n")

	var reports [][]int

	for _, s := range splited {
		splitedLevels := strings.Split(s, " ")

		var levels []int

		for _, sl := range splitedLevels {
			value, err := strconv.Atoi(sl)
			if err != nil {
				panic(err)
			}

			levels = append(levels, value)
		}

		reports = append(reports, levels)
	}

	return reports
}

func comparator(first int, second int) int {
	if first == second {
		return 0
	}

	if first < second {
		return -1
	}

	return 1
}

func hasError(order int, left int, right int) bool {
	if comparator(left, right) != order {
		return true
	}

	if math.Abs(float64(left-right)) > 3 {
		return true
	}

	return false
}

func checkReport(report []int) int {
	order := comparator(report[0], report[1])

	if order == 0 {
		if hasError(order, report[1], report[2]) {
			return 1
		}

		return 0
	}

	for i := 0; i < len(report)-1; i++ {
		if hasError(order, report[i], report[i + 1]) {
			if i + 2 >= len(report) {
				return i + 1
			}

			if hasError(order, report[i + 1], report[i + 2]) {
				return i + 1
			}

			return i
		}
	}

	return -1
}

func part1(reports [][]int) {
	res := 0

	for _, report := range reports {
		safe := 1

		if len(report) == 0 {
			continue
		}

		order := comparator(report[0], report[1])

		if order == 0 {
			continue
		}

		for i := 0; i < len(report)-1; i++ {
			if comparator(report[i], report[i+1]) != order {
				safe = 0
				break
			}

			if math.Abs(float64(report[i]-report[i+1])) > 3 {
				safe = 0
				break
			}
		}

		res += safe
	}

	fmt.Printf("Part 1 : %d\n", res)
}

func part2(reports [][]int) {
	res := 0

	for _, report := range reports {
		if len(report) == 0 {
			continue
		}

		ret := checkReport(report)
		if ret != -1 {
			fmt.Println("BASE: ")
			fmt.Println(report)
			newReport := append(report[:ret], report[ret+1:]...)
			fmt.Println("NEW: ")
			fmt.Println(newReport)
			ret = checkReport(newReport)

			fmt.Printf("OK: %t\n", ret == -1)
			fmt.Println("---------------------------")
		}

		if ret == -1 {
			res++
		}
	}

	fmt.Printf("Part 2 : %d\n", res)
}

func main() {
	content := readFile()

	reports := parseInput(content)

	part1(reports)
	part2(reports)
}
