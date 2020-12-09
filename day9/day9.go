package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func getInput() []int {
	var inputs []int
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if err != nil {
			fmt.Println("could not parse", err)
		}
		data, _ := strconv.Atoi(scanner.Text())
		inputs = append(inputs, data)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func findAllSumPairs(numbers []int) []int {
	sums := make([]int, 0)
	for index, number := range numbers {
		if index == len(numbers) {
			return sums
		}
		for jndex, secondNumber := range numbers {
			if index == jndex {
				continue
			}
			sum := number + secondNumber
			sums = append(sums, sum)
		}
	}
	return sums
}

func main() {
	defer elapsed("AoC Day 9")()
	numbers := getInput()
	answer1 := part1(numbers, 25)
	answer2 := part2(numbers, answer1)
	fmt.Println("A1:", answer1, "\nA2:", answer2)
}

func contains(numbers []int, number int) bool {
	for _, num := range numbers {
		if num == number {
			return true
		}
	}
	return false
}

func part1(numbers []int, preamble int) int {
	defer elapsed("Part 1")()
	i := preamble
	for i < len(numbers) {
		numberWindow := numbers[i-preamble : i]
		sums := findAllSumPairs(numberWindow)
		if !contains(sums, numbers[i]) {
			return numbers[i]
		}
		i++
	}
	return 0
}

func part2(numbers []int, target int) int {
	defer elapsed("Part 2")()
	for index, _ := range numbers {
		numberSeries := make([]int, 0)
		j := index
		sum := 0
		for j < len(numbers)-index {
			numberSeries = append(numberSeries, numbers[j])
			sum += numbers[j]
			if sum == target {
				sort.Ints(numberSeries)
				return numberSeries[0] + numberSeries[len(numberSeries)-1]
				break
			}
			j++
		}
	}
	return 0
}
