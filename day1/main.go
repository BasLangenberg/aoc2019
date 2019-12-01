package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var result int

	f, err := os.Open("input.txt")
	if err != nil {
		fmt.Errorf("Unable to open file")
	}

	reader := bufio.NewReader(f)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		input, err := strconv.Atoi(string(line))
		if err != nil {
			fmt.Errorf("Unable to parse input %s", line)
		}

		mod := moduleFuel(input)
		fuel := fuelFuel(mod)
		result += mod + fuel
	}

	fmt.Printf("Total fuel for modules and fuel: %d", result)
}

// Te laag: 1681920
// Te laag: 1681921
// Te laag: 1681921
// NOK:     5045850
// NOK:     4485007

func fuelFuel(i int) int {
	var fuelAmount int

	for {
		if i <= 0 {
			break
		}

		i = (i / 3) - 2
		if i > 0 {
			fuelAmount += i
		}
	}

	return fuelAmount
}

func moduleFuel(num int) int {
	return (num / 3) - 2
}
