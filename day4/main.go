package main

import (
	"fmt"
	"strconv"
)

func main() {
	var count1 int
	var count2 int
	for index := 168630; index <= 718098; index++ {
		if validate(index) {
			count1++
		}
	}

	for index := 168630; index <= 718098; index++ {
		if validate2(index) {
			count2++
		}
	}

	fmt.Println(count1)
	fmt.Println(count2)
}

func realPair(input string) bool {
	set := make(map[string]int)
	for index := 0; index < len(input)-1; index++ {
		if input[index] == input[index+1] {
			if set[string(input[index])] == 0 {
				set[string(input[index])]++
			}
			set[string(input[index])]++
		}
	}
	fmt.Println(set)
	for _, num := range set {
		if num == 2 {
			return true
		}
	}
	return false
}

func validate2(input int) bool {
	if !validate(input) {
		return false
	}

	if !realPair(strconv.Itoa(input)) {
		return false
	}

	return true
}

func validate(input int) bool {
	sinput := strconv.Itoa(input)
	if !checkDup(sinput) {
		return false
	}
	if !checkIncr(sinput) {
		return false
	}

	return true
}

func checkDup(input string) bool {
	for index := 0; index < len(input)-1; index++ {
		if input[index] == input[index+1] {
			return true
		}
	}
	return false
}

func checkIncr(input string) bool {
	for index := 0; index < len(input)-1; index++ {
		a, _ := strconv.Atoi(string(input[index]))
		b, _ := strconv.Atoi(string(input[index+1]))

		if a > b {
			return false
		}
	}
	return true
}
