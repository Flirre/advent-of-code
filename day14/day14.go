package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	defer elapsed("AoC Day 14")()
	input := getInput()
	part1(input)
	// part2(input)
}

func part1(input []string) {
	defer elapsed("Part 1")()
	var mem [100000]int64
	zeroes := int64(0)
	ones := int64(0)
	sum := int64(0)
	for _, instruction := range input {
		ParseInstruction(instruction, &mem, &zeroes, &ones)
	}
	for _, memValue := range mem {
		sum += memValue
	}
	fmt.Println(sum)
}

func ParseInstruction(rawInstruction string, mem *[100000]int64, zeroes *int64, ones *int64) {
	splitInstruction := strings.Split(rawInstruction, "=")
	leftSide := splitInstruction[0]
	rightSide := splitInstruction[1]
	numbers := regexp.MustCompile(`\d+`)
	if strings.Contains(leftSide, "mask") {
		*zeroes = 0
		*ones = 0
		UpdateMask(rightSide, zeroes, ones)
	}
	if strings.Contains(leftSide, "mem") {
		memPosString := numbers.FindString(leftSide)
		memPos, _ := strconv.ParseInt(memPosString, 10, 64)
		value := numbers.FindString(rightSide)
		_, intValue := StringToBit(value)
		masked := intValue & (137438953471 - *zeroes)
		masked = (masked | *ones)
		mem[memPos] = masked

	}
}

func StringToBit(number string) (string, int64) {
	intValue, _ := strconv.ParseInt(number, 10, 64)
	bitValue := strconv.FormatInt(intValue, 2)
	return bitValue, intValue
}

func UpdateMask(newMask string, zeroes *int64, ones *int64) {
	for index, bit := range newMask {
		if string(bit) == "0" {
			*zeroes += powInt(2, 36-index)
		}
		if string(bit) == "1" {
			*ones += powInt(2, 36-index)
		}
	}

}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func powInt(x, y int) int64 {
	return int64(math.Pow(float64(x), float64(y)))
}

// func part2(input []string) {
// 	defer elapsed("Part 2")()

// 	var mem [100000]int64
// 	zeroes := int64(0)
// 	ones := int64(0)
// 	sum := int64(0)
// 	for _, instruction := range input {
// 		ParseInstruction(instruction, &mem, &zeroes, &ones)
// 	}
// 	for _, memValue := range mem {
// 		sum += memValue
// 	}
// 	fmt.Println(sum)
// }
