package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"
)

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

func part1(seatMap []string) int {
	defer elapsed("Part 1")()

	currentSeatMap := seatMap
	changed := true
	for changed {
		currentSeatMap, changed = round1(currentSeatMap)
	}

	occupados := 0
	for _, row := range currentSeatMap {
		for _, chair := range row {
			if string(chair) == "#" {
				occupados += 1
			}
		}
	}
	return occupados
}

func round1(seatMap []string) ([]string, bool) {
	newSeatMap := make([]string, len(seatMap))
	for i, row := range seatMap {
		newRow := make([]rune, len(row))
		for j, placement := range []rune(row) {
			switch string(placement) {
			case "L":
				newRow[j] = switchChar('L', occupado(i, j, seatMap) == 0)
			case "#":
				newRow[j] = switchChar('#', occupado(i, j, seatMap) >= 4)
			default:
				newRow[j] = placement
			}
		}
		newSeatMap[i] = string(newRow)
	}
	return newSeatMap, !reflect.DeepEqual(seatMap, newSeatMap)
}

func part2(seatMap []string) int {
	defer elapsed("Part 2")()

	currentSeatMap := seatMap
	changed := true
	for changed {
		currentSeatMap, changed = round2(currentSeatMap)
	}

	occupados := 0
	for _, row := range currentSeatMap {
		for _, chair := range row {
			if string(chair) == "#" {
				occupados += 1
			}
		}
	}
	return occupados
}

func round2(seatMap []string) ([]string, bool) {
	newSeatMap := make([]string, len(seatMap))
	for i, row := range seatMap {
		newRow := make([]rune, len(row))
		for j, placement := range []rune(row) {
			switch string(placement) {
			case "L":
				newRow[j] = switchChar('L', occupado2(i, j, seatMap) == 0)
			case "#":
				newRow[j] = switchChar('#', occupado2(i, j, seatMap) >= 5)
			default:
				newRow[j] = placement
			}
		}
		newSeatMap[i] = string(newRow)
	}
	return newSeatMap, !reflect.DeepEqual(seatMap, newSeatMap)
}

func switchChar(char rune, condition bool) rune {
	if !condition {
		return char
	}
	if char == 76 {
		return 35
	}
	if char == 35 {
		return 76
	}
	return char
}

func occupado(y int, x int, seatMap []string) int {
	occupiedChairs := 0
	for i := (y - 1); i <= (y + 1); i++ {
		for j := (x - 1); j <= (x + 1); j++ {
			if i < 0 || i > len(seatMap)-1 || j < 0 || j > len(seatMap[0])-1 || (i == y && j == x) {
				continue
			}
			if string(seatMap[i][j]) == "#" {
				occupiedChairs++
			}
		}
	}
	return occupiedChairs
}

func occupado2(y int, x int, seatMap []string) int {
	occupiedChairs := 0
	//top
	for i := y - 1; i >= 0; i-- {
		if string(seatMap[i][x]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][x]) == "#" || string(seatMap[i][x]) == "L" {
			break
		}
	}
	//top-right
	for i, j := y-1, x+1; i >= 0 && (j < len(seatMap[0])); i, j = i-1, j+1 {
		if string(seatMap[i][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][j]) == "#" || string(seatMap[i][j]) == "L" {
			break
		}
	}
	//right
	for j := x + 1; j < len(seatMap[0]); j++ {
		if string(seatMap[y][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[y][j]) == "#" || string(seatMap[y][j]) == "L" {
			break
		}
	}
	//bottom-right
	for i, j := y+1, x+1; i < len(seatMap) && (j < len(seatMap[0])); i, j = i+1, j+1 {
		if string(seatMap[i][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][j]) == "#" || string(seatMap[i][j]) == "L" {
			break
		}
	}
	//bottom
	for i := y + 1; i < len(seatMap); i++ {
		if string(seatMap[i][x]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][x]) == "#" || string(seatMap[i][x]) == "L" {
			break
		}
	}
	//bottom-left
	for i, j := y+1, x-1; i < len(seatMap) && j >= 0; i, j = i+1, j-1 {
		if string(seatMap[i][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][j]) == "#" || string(seatMap[i][j]) == "L" {
			break
		}
	}
	//left
	for j := x - 1; j >= 0; j-- {
		if string(seatMap[y][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[y][j]) == "#" || string(seatMap[y][j]) == "L" {
			break
		}
	}
	//top-left
	for i, j := y-1, x-1; i >= 0 && (j >= 0); i, j = i-1, j-1 {
		if string(seatMap[i][j]) == "#" {
			occupiedChairs++
		}
		if string(seatMap[i][j]) == "#" || string(seatMap[i][j]) == "L" {
			break
		}
	}
	return occupiedChairs
}
