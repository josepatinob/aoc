package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums, err := readFile("day1/input.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}

	// part 1
	// for i := 0; i < len(nums)-1; i++ {
	// 	for j := i + 1; j < len(nums); j++ {
	// 		if nums[i]+nums[j] == 2020 {
	// 			result := nums[i] * nums[j]
	// 			fmt.Printf("%d \n", result)
	// 			break
	// 		}
	// 	}
	// }

	// part 2
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for x := j + 1; x < len(nums); x++ {
				if nums[i]+nums[j]+nums[x] == 2020 {
					result := nums[i] * nums[j] * nums[x]
					fmt.Printf("%d \n", result)
					break
				}
			}
		}
	}
}

func readFile(fName string) (nums []int, err error) {
	data, err := ioutil.ReadFile(fName)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil {
			return nil, err
		}
		nums = append(nums, n)
	}

	return nums, nil
}
