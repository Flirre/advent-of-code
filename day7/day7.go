package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

// lack of negative lookbehind makes this regex solution less elegant
func main() {
	bags := make(map[string][]string)
	bags2 := make(map[string][]string)
	bagsAr := make([]string, 0)
	bagsAr2 := make([]string, 0)
	var pattern = regexp.MustCompile(`([a-z]+\s\w+\S|\d\s\w+\s\w+)`)
	var colorOnly = regexp.MustCompile(`([a-z]+\s[a-z]+)`)
	var nrAndName = regexp.MustCompile(`(\d+\s\w+\s\w+)`)
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
		}
	}
	for _, bag2 := range input {
		matches2 := pattern.FindAllString(bag2, -1)
		parentBag2 := matches2[0]
		for _, name2 := range matches2 {
			trimmedBagName2 := colorOnly.FindAllString(name2, -1)
			trimmedBagName3 := nrAndName.FindAllString(name2, -1)

			if parentBag2 == trimmedBagName2[0] || "bags contain" == trimmedBagName2[0] {
				continue
			}
			if "no other" == trimmedBagName2[0] {
				bags2[parentBag2] = append(bags2[parentBag2], "1")
			} else {
				bags2[parentBag2] = append(bags2[parentBag2], trimmedBagName3[0])
			}
		}
	}
	b := bagToTop("shiny gold", bags, &bagsAr)
	c := removeDuplicates(*b)
	fmt.Println("Answer part 1:", len(c)-1)
	d := bagToBottom("shiny gold", bags2, &bagsAr2, 1)
	fmt.Println("Answer part 2:", len(*d)-1)
}

func bagToTop(bagPtrn string, bags map[string][]string, bagsAr *[]string) *[]string {
	*bagsAr = append(*bagsAr, bagPtrn)
	for _, bag := range bags[bagPtrn] {
		bagToTop(bag, bags, bagsAr)
		bags[bag] = nil
	}
	return bagsAr
}

func bagToBottom(bagPtrn string, bags map[string][]string, bagsAr *[]string, count int) *[]string {
	var colorOnly = regexp.MustCompile(`([a-z]+\s[a-z]+)`)
	var nrOnly = regexp.MustCompile(`\d+`)

	for i := 0; i < count; i++ {
		*bagsAr = append(*bagsAr, bagPtrn)
	}
	for _, bag := range bags[bagPtrn] {
		if bag == "1" {
			return bagsAr
		}
		newBag := colorOnly.FindAllString(bag, -1)
		bagNum := nrOnly.FindAllString(bag, -1)
		bagNr, _ := strconv.Atoi(bagNum[0])
		for i := 0; i < count; i++ {
			bagToBottom(newBag[0], bags, bagsAr, bagNr)
		}
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
