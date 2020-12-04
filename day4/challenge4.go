package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// part 1
func main() {
	data, err := ioutil.ReadFile("day4/input.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	lines := strings.Split(string(data), "\n\n")
	requiredFields := [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	validPassports := 0

	for i := 0; i < len(lines); i++ {
		invalidPassport := false
		for j := 0; j < len(requiredFields); j++ {
			index := strings.Index(lines[i], requiredFields[j])
			correctLine := lines[i]
			if index == -1 {
				invalidPassport = true
			} else {
				isValid := checkValidity(correctLine[index:], requiredFields[j])
				if !isValid {
					invalidPassport = true
				}
			}
		}
		if !invalidPassport {
			validPassports++
		}
	}

	fmt.Println("Valid: ", validPassports)
}

// part 2
func checkValidity(passportAttr string, attrCode string) (isValid bool) {
	spaceIndex := strings.Index(passportAttr, " ")
	newLineIndex := strings.Index(passportAttr, "\n")
	indexedString := ""

	/** due to implementing this using strings only it caused a big headache
	**	when the lines ended in different ways
	 */
	if spaceIndex != -1 && newLineIndex != -1 {
		indexedString = passportAttr[0:minimum(spaceIndex, newLineIndex)]
	} else if spaceIndex == -1 && newLineIndex != -1 {
		indexedString = passportAttr[0:newLineIndex]
	} else if spaceIndex != -1 && newLineIndex == -1 {
		indexedString = passportAttr[0:spaceIndex]
	} else {
		indexedString = passportAttr[0:]
	}

	colonIndex := strings.Index(passportAttr, ":") + 1

	value := strings.TrimSpace(indexedString[colonIndex:])

	switch attrCode {
	case "byr":
		//fmt.Println("byr", value)
		year, _ := strconv.Atoi(value)
		//fmt.Println("year: ", year)
		return year >= 1920 && year <= 2002
	case "iyr":
		//fmt.Println("iyr", value)
		issued, _ := strconv.Atoi(value)
		return issued >= 2010 && issued <= 2020
	case "eyr":
		//fmt.Println("eyr", value)
		expires, _ := strconv.Atoi(value)
		return expires >= 2020 && expires <= 2030
	case "hgt":
		//fmt.Println("hgt", value)
		return checkHeight(value)
	case "hcl":
		//fmt.Println("hcl", value)
		return checkHairColor(value)
	case "ecl":
		//fmt.Println("ecl", value)
		return checkEyeColor(value)
	case "pid":
		//fmt.Println("pid", value)
		return len(value) == 9
	default:
		return false
	}
}

func minimum(val1 int, val2 int) int {
	minVal := -1
	if val1 < val2 {
		minVal = val1
	} else {
		minVal = val2
	}
	return minVal
}

func checkEyeColor(str string) bool {
	requiredColors := [7]string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	contained := false
	for i := 0; i < len(requiredColors); i++ {
		if requiredColors[i] == str {
			contained = true
			break
		}
	}
	return contained
}

func checkHairColor(str string) bool {
	return string(str[0]) == string("#") && len(str) == 7
}

func checkHeight(str string) bool {
	validHeight := false
	cmIndex := strings.Index(str, "cm")
	inIndex := strings.Index(str, "in")

	if cmIndex != -1 {
		val, _ := strconv.Atoi(str[0:cmIndex])
		if val >= 150 && val <= 193 {
			validHeight = true
		}
	} else if inIndex != -1 {
		val, _ := strconv.Atoi(str[0:inIndex])
		if val >= 59 && val <= 76 {
			validHeight = true
		}
	}

	return validHeight
}
