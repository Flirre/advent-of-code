package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"time"
)

func mod(a, b int) int {
	m := a % b
	if a < 0 && b < 0 {
		m -= b
	}
	if a < 0 && b > 0 {
		m += b
	}
	return m
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

func getInput() []string {
	var inputs []string
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
		inputs = append(inputs, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func main() {
	defer elapsed("AoC Day 11")()
	input := getInput()
	answer1 := part1(input)
	fmt.Println(1, answer1)
	answer2 := part2(input)
	fmt.Println(2, answer2)
}

func part1(input []string) int {
	defer elapsed("Part 1")()
	direction := 0
	position := [2]int{0, 0}
	for _, row := range input {
		instruction, modifier := parseInstruction(row)
		executeInstruction(instruction, modifier, &direction, &position)
	}
	return Abs(position[0]) + Abs(position[1])
}

func parseInstruction(rawInstruction string) (string, int) {
	letters := regexp.MustCompile(`[A-Z]`)
	numbers := regexp.MustCompile(`\d+`)
	instruction := letters.FindString(rawInstruction)
	modifier, _ := strconv.Atoi(numbers.FindString(rawInstruction))
	return instruction, modifier
}

func executeInstruction(instruction string, value int, direction *int, position *[2]int) {
	directions := [4]string{"E", "S", "W", "N"}
	switch instruction {
	case "N":
		position[1] += value
	case "S":
		position[1] -= value
	case "E":
		position[0] += value
	case "W":
		position[0] -= value
	case "L":
		*direction = mod((*direction - (value / 90)), 4)
	case "R":
		*direction = mod((*direction + (value / 90)), 4)
	case "F":
		executeInstruction(directions[*direction], value, direction, position)
	}
}

func executeInstruction2(instruction string, value int, position *[2]int, waypoint *[2]int) {
	switch instruction {
	case "N":
		waypoint[1] += value
	case "S":
		waypoint[1] -= value
	case "E":
		waypoint[0] += value
	case "W":
		waypoint[0] -= value
	case "L":
		for i := 0; i < value/90; i++ {
			waypoint = rotateCounterClockwise(waypoint)
		}
	case "R":
		for i := 0; i < value/90; i++ {
			waypoint = rotateClockwise(waypoint)
		}
	case "F":
		position[0] += value * waypoint[0]
		position[1] += value * waypoint[1]
	}
}

func rotateCounterClockwise(waypoint *[2]int) *[2]int {
	tempWaypoint := make([]int, 2)
	tempWaypoint[0] = waypoint[0]
	tempWaypoint[1] = waypoint[1]
	waypoint[0] = -tempWaypoint[1]
	waypoint[1] = tempWaypoint[0]
	return waypoint
}

func rotateClockwise(waypoint *[2]int) *[2]int {
	tempWaypoint := make([]int, 2)
	tempWaypoint[0] = waypoint[0]
	tempWaypoint[1] = waypoint[1]
	waypoint[0] = tempWaypoint[1]
	waypoint[1] = -tempWaypoint[0]
	return waypoint
}

func part2(input []string) int {
	defer elapsed("Part 2")()
	position := [2]int{0, 0}
	waypoint := [2]int{10, 1}
	for _, row := range input {
		instruction, modifier := parseInstruction(row)
		executeInstruction2(instruction, modifier, &position, &waypoint)
	}
	return Abs(position[0]) + Abs(position[1])
}
