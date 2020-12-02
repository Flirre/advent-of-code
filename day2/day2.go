package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() [][]string {
	inputs := [][]string{}
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineInput := strings.Split(scanner.Text(), " ")
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

func main() {
	input := getInput()
	var validPasswordsOS []string
	var validPasswordsNS []string
	for _, password := range input {

		if checkValidityOldStyle(password) {
			validPasswordsOS = append(validPasswordsOS, password[2])
		}

		if checkValidityNewStyle(password) {
			validPasswordsNS = append(validPasswordsNS, password[2])

		}
	}
	fmt.Println("Amount of valid passwords according to the old style: ", len(validPasswordsOS))
	fmt.Println("Amount of valid passwords according to the new style: ", len(validPasswordsNS))
}

func checkValidityOldStyle(password []string) bool {
	lowerCountLimit, _ := strconv.Atoi(strings.Split(password[0], "-")[0])
	upperCountLimit, _ := strconv.Atoi(strings.Split(password[0], "-")[1])
	charToMatch := strings.TrimSuffix(password[1], ":")
	occurences := strings.Count(password[2], charToMatch)
	return occurences >= lowerCountLimit && occurences <= upperCountLimit
}

func checkValidityNewStyle(password []string) bool {
	firstPosition, _ := strconv.Atoi(strings.Split(password[0], "-")[0])
	secondPosition, _ := strconv.Atoi(strings.Split(password[0], "-")[1])
	charToMatch := strings.TrimSuffix(password[1], ":")
	var firstMatch bool = string(password[2][firstPosition-1]) == charToMatch
	var secondMatch bool = string(password[2][secondPosition-1]) == charToMatch
	return ((firstMatch && secondMatch) == false) && (firstMatch || secondMatch)
}
