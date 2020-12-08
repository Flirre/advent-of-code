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

	reversedInstructions := make(map[string]string)
	reversedInstructions["nop"] = "jmp"
	reversedInstructions["jmp"] = "nop"
	reversedInstructions["acc"] = "acc"

	PC := 0
	accumulator := 0
	for PC < len(instructions) {
		instruction, modifier := parseInstruction(instructions[PC])
		if executedInstructions[PC] != "" {
			break
		}
		executedInstructions[PC] = instructions[PC]
		switch instruction {
		case "nop":
			PC += 1
		case "acc":
			PC += 1
			accumulator += modifier
		case "jmp":
			PC += modifier
		}
	}
	fmt.Println("\nanswer p1", accumulator)
}

func parseInstruction(instr string) (string, int) {
	splitInstruction := strings.Split(instr, " ")
	instruction := splitInstruction[0]
	modifier, _ := strconv.Atoi(splitInstruction[1])
	return instruction, modifier
}
