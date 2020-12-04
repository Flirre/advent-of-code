package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getInput() []map[string]string {
	var inputs []map[string]string
	passport := make(map[string]string)
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
			inputs = append(inputs, passport)
			passport = make(map[string]string)
		} else {
			pairs := strings.Split(scanner.Text(), " ")
			for _, pair := range pairs {
				keyValue := strings.Split(pair, ":")
				passport[keyValue[0]] = keyValue[1]
			}
		}
	}
	inputs = append(inputs, passport)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return inputs
}

func main() {
	input := getInput()
	validPassports := 0
	actuallyValidPassports := 0
	for _, passport := range input {
		if isValidPassport(passport, true) {
			validPassports += 1
		}
		if isValidPassport(passport, false) {
			actuallyValidPassports += 1
		}
	}
	fmt.Println(validPassports)
	fmt.Println(actuallyValidPassports)
}

func extractKeys(mapped map[string]string) []string {
	i := 0
	keys := make([]string, len(mapped))
	for k := range mapped {
		keys[i] = k
		i++
	}
	return keys
}

func isValidPassport(passport map[string]string, firstStar bool) bool {
	requiredKeys := []string{"ecl", "pid", "eyr", "hcl", "byr", "iyr", "hgt"}
	keys := extractKeys(passport)
	if len(keys) < len(requiredKeys) {
		return false
	}
	for _, requiredKey := range requiredKeys {
		if !contains(keys, requiredKey) {
			return false
		}
	}
	if firstStar {
		return true
	}
	return isValidBYR(passport["byr"]) && isValidIYR(passport["iyr"]) && isValidEYR(passport["eyr"]) && isValidHGT(passport["hgt"]) && isValidHCL(passport["hcl"]) && isValidECL(passport["ecl"]) && isValidPID(passport["pid"])
}

func isValidBYR(byr string) bool {
	matched, _ := regexp.MatchString(`^(19[2-9][0-9]|200[0-2])$`, byr)
	return matched
}

func isValidIYR(iyr string) bool {
	matched, _ := regexp.MatchString(`^(201[0-9]|2020)$`, iyr)
	return matched
}

func isValidEYR(eyr string) bool {
	matched, _ := regexp.MatchString(`^(202[0-9]|2030)$`, eyr)
	return matched
}

func isValidHGT(hgt string) bool {
	matched, _ := regexp.MatchString(`^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)$`, hgt)
	return matched
}

func isValidHCL(hcl string) bool {
	matched, _ := regexp.MatchString(`^(#)([0-9]|[a-f]){6}$`, hcl)
	return matched
}

func isValidECL(ecl string) bool {
	matched, _ := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, ecl)
	return matched
}

func isValidPID(pid string) bool {
	matched, _ := regexp.MatchString(`^([0-9]){9}$`, pid)
	return matched
}
