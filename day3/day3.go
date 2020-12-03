package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

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
	x0 := 0
	x1 := 0
	x2 := 0
	x3 := 0
	x4 := 0
	treeCount0 := 0
	treeCount1 := 0
	treeCount2 := 0
	treeCount3 := 0
	treeCount4 := 0
	for index, line := range input {
		tobogganRiding(1, 1, line, index, &treeCount0, &x0)
		tobogganRiding(3, 1, line, index, &treeCount1, &x1)
		tobogganRiding(5, 1, line, index, &treeCount2, &x2)
		tobogganRiding(7, 1, line, index, &treeCount3, &x3)
		tobogganRiding(1, 2, line, index, &treeCount4, &x4)
	}
	fmt.Println(treeCount1)
	fmt.Println(treeCount0 * treeCount1 * treeCount2 * treeCount3 * treeCount4)
}

func isTree(mapLine string, position int) bool {
	return string(mapLine[position%len(mapLine)]) == "#"
}

func tobogganRiding(right int, down int, line string, index int, treeCount *int, x *int) {
	if index == 0 {
		return
	}
	if down == 2 {
		if index%2 == 1 {
			return
		}
	}
	*x += right
	if isTree(line, *x) {
		*treeCount += 1
	}
}
