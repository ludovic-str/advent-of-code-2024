package main

import (
	"fmt"
	"os"
	"regexp"
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

func part1(content string) {
	r, _ := regexp.Compile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches := r.FindAllStringSubmatch(content, -1)

	res := 0

	for _, match := range matches {
		first, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}

		second, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}

		res += first * second
	}

	fmt.Printf("Part 1: %d\n", res)
}

func part2(content string) {
	r, _ := regexp.Compile("(mul\\((\\d{1,3}),(\\d{1,3})\\))|(do\\(\\))|(don't\\(\\))")

	matches := r.FindAllStringSubmatch(content, -1)

	disabled := false

	res := 0

	for _, match := range matches {
		if match[0] == "do()" {
			disabled = false
		} else if match[0] == "don't()" {
			disabled = true
		} else if strings.Contains(match[0], "mul") && !disabled {
			first, err := strconv.Atoi(match[2])
			if err != nil {
				fmt.Println(match)
				panic(err)
			}

			second, err := strconv.Atoi(match[3])
			if err != nil {
				panic(err)
			}

			res += first * second
		}
	}

	fmt.Printf("Part 2: %d\n", res)
}

func main() {
	content := readFile()

	part1(content)
	part2(content)
}
