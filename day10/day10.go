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
	sort.Ints(inputs)
	inputs = append(inputs, inputs[len(inputs)-1]+3)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func main() {
	defer elapsed("AoC Day 10")()
	input := getInput()
	answer1 := part1(input)
	fmt.Println(answer1)
	answer2 := part2(input, true)
	fmt.Println(answer2)
}

func contains(numbers []int, number int) bool {
	for _, num := range numbers {
		if num == number {
			return true
		}
	}
	return false
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func part1(joltages []int) int {
	defer elapsed("Part 1")()
	prevJoltage := 0
	joltageDifferences := make(map[int]int)
	for _, joltage := range joltages {
		diff := joltage - prevJoltage
		prevJoltage = joltage
		joltageDifferences[diff] += 1
	}
	return joltageDifferences[1] * joltageDifferences[3]
}

func part2(joltages []int, first bool) int {
	defer elapsed("Part 2")()
	if first {
		joltages = append([]int{0}, joltages...)
	}

	combos := make([]int, len(joltages))
	combos[0] = 1
	combos[1] = 1

	for i := 2; i < len(joltages); i++ {
		for j := i - 1; j >= 0; j-- {
			if joltages[i]-joltages[j] <= 3 {
				combos[i] += combos[j]
			} else {
				break
			}
		}
	}
	return combos[len(combos)-1]
}
