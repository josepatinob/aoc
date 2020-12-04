package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// part 2
func main() {
	value := traverseRoad(1, 1) * traverseRoad(3, 1) * traverseRoad(5, 1) * traverseRoad(7, 1) * traverseRoad(1, 2)
	fmt.Println("Total: ", value)
}

func concatStrings(str string, times int) (val string) {
	for i := 0; i < times+1; i++ {
		val += str
	}

	return val
}

// part 1
func traverseRoad(r int, d int) (numTrees int) {
	totalTrees := 0
	data, err := ioutil.ReadFile("day3/input.txt")
	lines := strings.Split(string(data), "\n")
	verticleSize := len(lines)
	horizontalSizeMultiple := verticleSize / 2
	right := r
	road := make([]string, 0, verticleSize)

	// fmt.Println(verticleSize)
	// fmt.Println(horizontalSize)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < verticleSize; i++ {
		road = append(road, concatStrings(lines[i], horizontalSizeMultiple))
	}

	for j := d; j < verticleSize; j = j + d {
		value := road[j]
		// fmt.Println(string(value[right]))
		// fmt.Println(string("#"))
		if string(value[right]) == string("#") {
			totalTrees++
		}
		right += r
	}

	return totalTrees
}
