package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("day6/input.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	groups := strings.Split(string(data), "\n\n")
	totalYes := 0

	for i := 0; i < len(groups); i++ {
		group := groups[i]
		m := make(map[string]int)

		totalPeople := strings.Count(group, "\n") + 1

		for j := 0; j < len(group); j++ {
			value := string(group[j])

			//fmt.Println(value)

			if string(value) != " " && string(value) != "\n" {
				_, found := m[value]
				if !found {
					m[value] = 1
				} else {
					v := m[value]
					m[value] = v + 1
				}
			}
		}

		fmt.Println(m)

		if totalPeople == 1 {
			for _, v := range m {
				totalYes += v
			}
		} else {
			for _, v := range m {
				if v == totalPeople {
					totalYes++
				}
			}
		}
	}

	fmt.Println("total: ", totalYes)
}
