package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,13,1,19,1,5,19,23,2,10,23,27,1,27,5,31,2,9,31,35,1,35,5,39,2,6,39,43,1,43,5,47,2,47,10,51,2,51,6,55,1,5,55,59,2,10,59,63,1,63,6,67,2,67,6,71,1,71,5,75,1,13,75,79,1,6,79,83,2,83,13,87,1,87,6,91,1,10,91,95,1,95,9,99,2,99,13,103,1,103,6,107,2,107,6,111,1,111,2,115,1,115,13,0,99,2,0,14,0"

func getInput() []int {
	var data []int
	for _, num := range strings.Split(input, ",") {
		conv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Failed to convert num: %s", num)
		}
		data = append(data, conv)
	}

	return data
}

// Write a test!
func main() {
	data := getInput()
	data[1] = 12
	data[2] = 2
	compute(data)
	fmt.Println(data[0])

	output := 19690720
	for index := 0; index < 100; index++ {
		for index2 := 0; index2 < 100; index2++ {
			data = getInput()
			data[1] = index
			data[2] = index2
			compute(data)
			if data[0] == output {
				fmt.Printf("%d and %d\n", index, index2)
				break
			}
		}
	}
}

func compute(input []int) int {
	stepper := 0
	for {
		if stepper >= len(input)-1 {
			break
		}
		s1 := input[stepper+1]
		s2 := input[stepper+2]
		s3 := input[stepper+3]

		switch input[stepper] {
		case 1:
			input[s3] = input[s1] + input[s2]
		case 2:
			input[s3] = input[s1] * input[s2]
		case 99:
			break

		}

		stepper += 4

	}

	return input[0]
}
