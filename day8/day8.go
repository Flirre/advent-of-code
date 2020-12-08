package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	instructions := getInput()
	executedInstructions := make(map[int]string)
	i := 0
	accumulator := 0
	for i < len(instructions) {
		if executedInstructions[i] != "" {
			fmt.Println(i)
			break
		}
		// fmt.Println(i, instructions[i])
		splitInstruction := strings.Split(instructions[i], " ")
		executedInstructions[i] = instructions[i]
		instruction := splitInstruction[0]
		modifier, _ := strconv.Atoi(splitInstruction[1])
		switch instruction {
		case "nop":
			//fmt.Println("NOP", i)
			i += 1
		case "acc":
			i += 1
			accumulator += modifier
		case "jmp":
			//fmt.Println("JMP", i)
			i += modifier
		}
	}
	fmt.Println("\n\n\nanswer p1", accumulator)
}
