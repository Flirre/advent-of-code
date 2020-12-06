package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func getInput() [][]string {
	var inputs [][]string
	group := make([]string, 0)
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
		if len(scanner.Text()) == 0 {
			inputs = append(inputs, group)
			group = make([]string, 0)
		} else {
			group = append(group, scanner.Text())
		}
	}
	inputs = append(inputs, group)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func main() {
	input := getInput()
	sumAnyone := 0
	sumEveryone := 0
	for _, group := range input {
		sumAnyone = sumAnyone + len(removeDuplicates(group))
		sumEveryone = sumEveryone + commonChars(group)
	}
	fmt.Println("Answer part 1:", sumAnyone)
	fmt.Println("Answer part 2:", sumEveryone)
}

func removeDuplicates(a []string) []string {
	check := make(map[string]int)
	res := make([]string, 0)
	for _, val := range a {
		if len(val) > 1 {
			for _, subval := range val {
				check[string(subval)] = 1
			}
		} else {
			check[val] = 1
		}
	}

	for letter, _ := range check {
		res = append(res, letter)
	}
	return res
}

// following func shamefully borrowed from https://github.com/halfrost/LeetCode-Go/blob/master/leetcode/1002.Find-Common-Characters/1002.%20Find%20Common%20Characters.go
// TODO: unborrow this func and stop bringing disgrace upon my family.
func commonChars(A []string) int {
	cnt := [26]int{}
	for i := range cnt {
		cnt[i] = math.MaxUint16
	}
	cntInWord := [26]int{}
	for _, word := range A {
		for _, char := range []byte(word) { // compiler trick - here we will not allocate new memory
			cntInWord[char-'a']++
		}
		for i := 0; i < 26; i++ {

			if cntInWord[i] < cnt[i] {
				cnt[i] = cntInWord[i]
			}
		}

		for i := range cntInWord {
			cntInWord[i] = 0
		}
	}
	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < cnt[i]; j++ {
			result = append(result, string(rune(i+'a')))
		}
	}
	return len(result)
}
