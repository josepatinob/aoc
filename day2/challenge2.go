package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	totalGoodPwds := 0
	data, err := ioutil.ReadFile("day2/input.txt")
	passwords := strings.Split(string(data), "\n")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// part 1
	// for i := 0; i < len(passwords); i++ {
	// 	min := getMin(passwords[i])
	// 	max := getMax(passwords[i])
	// 	character := getChar(passwords[i])
	// 	pwd := getPassword(passwords[i])
	// 	count := countInstances(pwd, character)

	// 	//fmt.Println(count)

	// 	if count >= min && count <= max {
	// 		totalGoodPwds++
	// 	}
	// }

	// part 2
	for i := 0; i < len(passwords); i++ {
		position1 := getMin(passwords[i])
		position2 := getMax(passwords[i])
		character := getChar(passwords[i])
		pwd := getPassword(passwords[i])

		matchesPosition1 := false
		matchesPosition2 := false

		if strings.TrimSpace(string(pwd[position1])) == strings.TrimSpace(character) {
			matchesPosition1 = true
		}

		if strings.TrimSpace(string(pwd[position2])) == strings.TrimSpace(character) {
			matchesPosition2 = true
		}

		// fmt.Println(position1)
		// fmt.Println(position2)

		if (matchesPosition1 && !matchesPosition2) || (!matchesPosition1 && matchesPosition2) {
			totalGoodPwds++
		}

	}

	fmt.Println(totalGoodPwds)
}

func getMin(pass string) (min int) {
	dash := strings.Index(pass, "-")
	first := pass[0:dash]
	min, _ = strconv.Atoi(first)
	return min
}

func getMax(pass string) (max int) {
	dash := strings.Index(pass, "-")
	spaceIndex := strings.Index(pass, " ")
	second := pass[dash+1 : spaceIndex]
	max, _ = strconv.Atoi(second)
	return max
}

func getChar(pass string) (character string) {
	colon := strings.Index(pass, ":")
	spaceIndex := strings.Index(pass, " ")
	character = pass[spaceIndex:colon]
	//fmt.Println(character)
	return character
}

func getPassword(pass string) (password string) {
	colon := strings.Index(pass, ":")
	password = pass[colon+1:]
	//fmt.Println(password)
	return password
}

func countInstances(pass string, subString string) (count int) {
	for i := 0; i < len(pass); i++ {
		if strings.TrimSpace(string(pass[i])) == strings.TrimSpace(subString) {
			count++
		}
	}
	return count
}
