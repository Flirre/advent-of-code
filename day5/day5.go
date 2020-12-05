package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func getInput() []string {
	inputs := []string{}
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
	input := getInput()
	highestSeatID := 0
	var seatIDs []int
	for _, boardingPass := range input {
		_, _, seatID := binaryDecode(boardingPass)
		seatIDs = append(seatIDs, seatID)
		if seatID > highestSeatID {
			highestSeatID = seatID
		}
	}
	fmt.Println("Highest seat number:", highestSeatID)
	sort.Ints(seatIDs)
	for index, seat := range seatIDs {
		if seatIDs[index+1] == seat+2 {
			fmt.Println("My seat number:", seat+1)
			break
		}
	}
}

func binaryDecode(seatCode string) (int, int, int) {
	row := makeRange(0, 127)
	column := makeRange(0, 7)
	commands := strings.Split(seatCode, "")

	for index, command := range commands {
		if index < 7 {
			if command == "F" {
				row = row[0 : len(row)/2]
			}
			if command == "B" {
				row = row[len(row)/2:]
			}
		} else {
			if command == "L" {
				column = column[0 : len(column)/2]
			}
			if command == "R" {
				column = column[len(column)/2:]
			}
		}
	}
	seatID := row[0]*8 + column[0]

	return row[0], column[0], seatID
}
