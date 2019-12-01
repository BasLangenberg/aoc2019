package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("Unable to open file")
	}
	i := day1(f)
	fuelforfuel(i)
}

// Te laag: 1681920
// Te laag: 1681921
// Te laag: 1681921
// NOK:     5045850

func fuelforfuel(i int) int {
	fmt.Println(i)

	mass := i
	fuel := 0
	for {
		if mass <= 2 {
			break
		}

		mass = (mass / 3) - 2
		if mass > 0 {
			fuel += mass
		}
		fmt.Println(mass)
	}

	fmt.Println(fuel + i)
	return fuel + i
}

func day1(f *os.File) int {

	reader := bufio.NewReader(f)

	var result int

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		input, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Errorf("Unable to parse input %s", line)
		}

		result += (input / 3) - 2
	}

	fmt.Println(result)

	return result

}
