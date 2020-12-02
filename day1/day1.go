package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	secondStar := true
	input := getInput()
	for i, value := range input {
		for j, secondValue := range input {
			if j <= i {
				continue
			}
			if !secondStar {
				if value+secondValue == 2020 {
					calculateExpense(value, secondValue)
				}
			} else {
				for k, thirdValue := range input {
					if k <= i || k <= j {
						continue
					} else {
						if value+secondValue+thirdValue == 2020 {
							calculateMoreExpenses(value, secondValue, thirdValue)
						}
					}
				}
			}
		}
	}
}

func getInput() []int {
	inputs := []int{}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInput, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("could not parse", err)
		}
		inputs = append(inputs, lineInput)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func calculateExpense(first int, second int) {
	fmt.Println(first * second)
}

func calculateMoreExpenses(first int, second int, third int) {
	fmt.Println(first * second * third)
}
