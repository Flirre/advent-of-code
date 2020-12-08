package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n\n", what, time.Since(start))
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
	defer elapsed("AoC Day 8")()
	instructions := getInput()
	part1(instructions)
	part2(instructions)
}

func part1(instructions []string) {
	defer elapsed("Part 1")()

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
	fmt.Println("Part 1:", accumulator)
}

func part2(instructions []string) {
	defer elapsed("Part 2")()
	reversedInstructions := make(map[string]string)
	reversedInstructions["nop"] = "jmp"
	reversedInstructions["jmp"] = "nop"
	reversedInstructions["acc"] = "acc"

	i := 0
	for i < len(instructions) {
		loop := false
		executedInstructions := make(map[int]string)

		PC := 0
		accumulator := 0
		for PC < len(instructions) {
			instruction, modifier := parseInstruction(instructions[PC])
			if executedInstructions[PC] != "" {
				loop = true
				break
			}
			executedInstructions[PC] = instructions[PC]
			if PC == i {
				instruction = reversedInstructions[instruction]
			}

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
		if !loop {
			fmt.Println("Part 2:", accumulator)
		}
		i++
	}
}

func parseInstruction(instr string) (string, int) {
	splitInstruction := strings.Split(instr, " ")
	instruction := splitInstruction[0]
	modifier, _ := strconv.Atoi(splitInstruction[1])
	return instruction, modifier
}
