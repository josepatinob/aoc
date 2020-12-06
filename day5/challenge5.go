package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day5/input.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	passports := strings.Split(string(data), "\n")

	fmt.Println(len(passports))

	seatIDs := makeSeatSlice(len(passports))

	for i := 0; i < len(passports); i++ {
		passport := passports[i]
		arr := makeArray()
		columns := makeColumnSlice()

		for j := 0; j < len(passport)-2; j++ {
			if string(passport[j]) == string("F") {
				mid := int(len(arr) / 2)
				arr = arr[0:mid]
			} else if string(passport[j]) == string("B") {
				mid := int(len(arr) / 2)
				arr = arr[mid:]
				//fmt.Println(arr)
			}
		}

		for x := len(passport) - 3; x < len(passport); x++ {
			if string(passport[x]) == string("L") {
				mid := int(len(columns) / 2)
				columns = columns[0:mid]
			} else if string(passport[x]) == string("R") {
				mid := int(len(columns) / 2)
				columns = columns[mid:]
			}
		}

		row := arr[0]
		column := columns[0]
		seatID := row*8 + column

		seatIDs = append(seatIDs, seatID)
	}

	sort.Ints(seatIDs)

	mySeat := 0

	fmt.Println(len(seatIDs))

	for n := 1; n < len(seatIDs); n++ {
		if seatIDs[n]-seatIDs[n-1] == 2 {
			mySeat = seatIDs[n] - 1
		}
	}

	fmt.Println(mySeat)
}

func makeArray() []int {
	arr := make([]int, 128)
	for i := 0; i < 128; i++ {
		arr[i] = i
	}
	return arr
}

func makeColumnSlice() []int {
	columns := make([]int, 8)
	for i := 0; i < 8; i++ {
		columns[i] = i
	}
	return columns
}

func makeSeatSlice(size int) []int {
	seatIDs := make([]int, size)
	return seatIDs
}
