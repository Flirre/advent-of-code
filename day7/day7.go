package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

// ([a-z]+ \w+\S|\d+ \w+ \w+)\b(?<!bags|contain|contains)
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

// lack of negative lookbehind makes this regex solution less elegant
func main() {
	bags := make(map[string][]string)
	bagsAr := make([]string, 0)
	var pattern = regexp.MustCompile(`([a-z]+\s\w+\S|\d\s\w+\s\w+)`)
	var colorOnly = regexp.MustCompile(`([a-z]+\s[a-z]+)`)
	input := getInput()
	for _, bag := range input {
		matches := pattern.FindAllString(bag, -1)
		parentBag := matches[0]
		for _, name := range matches {
			trimmedBagName := colorOnly.FindAllString(name, -1)
			if "bags contain" == trimmedBagName[0] || "no other" == trimmedBagName[0] || parentBag == trimmedBagName[0] {
				continue
			}
			bags[trimmedBagName[0]] = append(bags[trimmedBagName[0]], parentBag)
			//fmt.Println(parentBag, trimmedBagName)
		}
	}
	b := bagToTop("shiny gold", bags, &bagsAr)
	c := removeDuplicates(*b)
	fmt.Println("Answer part 1:", len(c)-1)
}

func bagToTop(bagPtrn string, bags map[string][]string, bagsAr *[]string) *[]string {
	*bagsAr = append(*bagsAr, bagPtrn)
	for _, bag := range bags[bagPtrn] {
		bagToTop(bag, bags, bagsAr)
		bags[bag] = nil
	}
	return bagsAr
}

func removeDuplicates(a []string) []string {
	check := make(map[string]int)
	res := make([]string, 0)
	for _, val := range a {
		check[val] = 1
	}

	for letter, _ := range check {
		res = append(res, letter)
	}
	return res
}
