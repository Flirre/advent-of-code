package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

func getInput() (int, []int, string) {
	var inputs string
	var time int
	var buses []int
	i := 0
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
		if i == 0 {
			fmt.Println(scanner.Text())
			time, _ = strconv.Atoi(scanner.Text())
		} else {
			inputs = scanner.Text()
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	numRegex := regexp.MustCompile(`\d+`)
	for _, num := range numRegex.FindAllString(inputs, -1) {
		numConv, _ := strconv.Atoi(num)
		buses = append(buses, numConv)
	}
	return time, buses, inputs
}

func main() {
	defer elapsed("AoC Day 11")()
	time, input1, input2 := getInput()
	fmt.Println(part1(time, input1))
	fmt.Println(part2(input2))
}

func part1(time int, buses []int) int {
	defer elapsed("Part 1")()
	shortestWait := 1000000000
	firstBus := 0
	for _, bus := range buses {
		newWait := bus - mod(time, bus)
		if newWait < shortestWait {
			shortestWait = newWait
			firstBus = bus
		}
	}
	return shortestWait * firstBus
}

func part2(buses string) int {
	defer elapsed("Part 2")()
	splitBuses := strings.Split(buses, ",")
	i := 0
	offset := 1
	correct := false
	for !correct {
		correct = true
		for index, bus := range splitBuses {
			if bus == "x" {
				continue
			}
			busInt, _ := strconv.Atoi(bus)

			if mod(i+index, busInt) != 0 {
				correct = false
				break
			}
			offset = offset * busInt
		}
		if !correct {
			i += offset
		}
		offset = 1
	}
	return i
}
