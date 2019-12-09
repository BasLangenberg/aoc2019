package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = "3,225,1,225,6,6,1100,1,238,225,104,0,1101,86,8,225,1101,82,69,225,101,36,65,224,1001,224,-106,224,4,224,1002,223,8,223,1001,224,5,224,1,223,224,223,102,52,148,224,101,-1144,224,224,4,224,1002,223,8,223,101,1,224,224,1,224,223,223,1102,70,45,225,1002,143,48,224,1001,224,-1344,224,4,224,102,8,223,223,101,7,224,224,1,223,224,223,1101,69,75,225,1001,18,85,224,1001,224,-154,224,4,224,102,8,223,223,101,2,224,224,1,224,223,223,1101,15,59,225,1102,67,42,224,101,-2814,224,224,4,224,1002,223,8,223,101,3,224,224,1,223,224,223,1101,28,63,225,1101,45,22,225,1101,90,16,225,2,152,92,224,1001,224,-1200,224,4,224,102,8,223,223,101,7,224,224,1,223,224,223,1101,45,28,224,1001,224,-73,224,4,224,1002,223,8,223,101,7,224,224,1,224,223,223,1,14,118,224,101,-67,224,224,4,224,1002,223,8,223,1001,224,2,224,1,223,224,223,4,223,99,0,0,0,677,0,0,0,0,0,0,0,0,0,0,0,1105,0,99999,1105,227,247,1105,1,99999,1005,227,99999,1005,0,256,1105,1,99999,1106,227,99999,1106,0,265,1105,1,99999,1006,0,99999,1006,227,274,1105,1,99999,1105,1,280,1105,1,99999,1,225,225,225,1101,294,0,0,105,1,0,1105,1,99999,1106,0,300,1105,1,99999,1,225,225,225,1101,314,0,0,106,0,0,1105,1,99999,7,677,677,224,102,2,223,223,1005,224,329,1001,223,1,223,1008,226,226,224,1002,223,2,223,1005,224,344,1001,223,1,223,1107,677,226,224,1002,223,2,223,1006,224,359,1001,223,1,223,107,677,677,224,102,2,223,223,1005,224,374,101,1,223,223,1108,677,226,224,102,2,223,223,1005,224,389,1001,223,1,223,1007,677,677,224,1002,223,2,223,1005,224,404,101,1,223,223,1008,677,226,224,102,2,223,223,1005,224,419,101,1,223,223,1108,226,677,224,102,2,223,223,1006,224,434,1001,223,1,223,8,677,226,224,1002,223,2,223,1005,224,449,101,1,223,223,1008,677,677,224,1002,223,2,223,1006,224,464,1001,223,1,223,1108,226,226,224,1002,223,2,223,1005,224,479,1001,223,1,223,1007,226,677,224,102,2,223,223,1005,224,494,1001,223,1,223,1007,226,226,224,102,2,223,223,1005,224,509,101,1,223,223,107,677,226,224,1002,223,2,223,1006,224,524,1001,223,1,223,108,677,677,224,102,2,223,223,1006,224,539,101,1,223,223,7,677,226,224,102,2,223,223,1006,224,554,1001,223,1,223,1107,226,677,224,102,2,223,223,1005,224,569,101,1,223,223,108,677,226,224,1002,223,2,223,1006,224,584,101,1,223,223,108,226,226,224,102,2,223,223,1006,224,599,1001,223,1,223,1107,226,226,224,102,2,223,223,1006,224,614,1001,223,1,223,8,226,677,224,102,2,223,223,1006,224,629,1001,223,1,223,107,226,226,224,102,2,223,223,1005,224,644,101,1,223,223,8,226,226,224,102,2,223,223,1006,224,659,101,1,223,223,7,226,677,224,102,2,223,223,1005,224,674,101,1,223,223,4,223,99,226"

// IntCode is the vm for this AOC problem
type IntCode struct {
	program                []int
	param1, param2, param3 bool
	pointer                int
	instruction            int
}

func main(){
	vm, _ := NewIntCode(input)
	vm.Run()
}

func (ic *IntCode) printOp() {
	fmt.Printf("Executing operation %d, pointer value %d", ic.instruction, ic.pointer)
}

// Run starts the program loaded in the IntCode VM
func (ic *IntCode) Run() {
	for {
		if ic.instruction == 99 {
			break
		}
		ic.ParseNextInstruction()
		switch ic.instruction {
		case 1:
			s1, s2, s3 := ic.parseMem()
			ic.program[s3] = s1 + s2
			ic.pointer += 4
		case 2:
			s1, s2, s3 := ic.parseMem()
			ic.program[s3] = s1 * s2
			ic.pointer += 4
		case 3:
			var intput int
			// TODO refactor to own function
			fmt.Print("Enter input: ")
			_, err := fmt.Scanf("%d", &intput)
			if err != nil {
				fmt.Println(err)
				panic("Unable to parse input int!")
			}
			// TODO Refactor parseMem to support this
			if ic.param1 {
				ic.program[ic.pointer+1] = intput
			} else {
				ic.program[ic.program[ic.pointer+1]] = intput
			}
			ic.pointer += 2
		case 4:
			var s1 int
			if ic.param1 {
				s1 = ic.program[ic.pointer+1]
			} else {
				s1 = ic.program[ic.program[ic.pointer+1]]
			}
			fmt.Printf("Output: %v\n", s1)
			ic.pointer += 2
		case 99:
			fmt.Println("Program exitting")
			break
		default:
			ic.printOp()
			fmt.Printf("Cant read instruction %d\n", ic.instruction)
			ic.program[ic.pointer] = 99
		}
	}
}

func (ic *IntCode) parseMem() (int, int, int) {
	fmt.Printf("Parsing memory for pointer location %d\n", ic.pointer)
	var s1, s2, s3, as1, as2, as3 int
	if ic.param1 {
		as1 = 0
		s1 = ic.program[ic.pointer+1]
	} else {
		as1 = ic.program[ic.pointer+1]
		s1 = ic.program[as1]
	}

	if ic.param2 {
		as2 = 0
		s2 = ic.program[ic.pointer+2]
	} else {
		as2 = ic.program[ic.pointer+2]
		s2 = ic.program[as2]
	}

	if ic.param3 {
		as3 = ic.program[ic.pointer+3]
		s3 = ic.program[as3]
	} else {
		as3 = 0
		s3 = ic.program[ic.pointer+3]
	}

	return s1, s2, s3
}

// NewIntCode does the setup of the IntCode VM and returns it to the caller
func NewIntCode(input string) (IntCode, error) {
	prog, _ := parseIntArrayFromString(input)

	ic := IntCode{
		program:     prog,
		param1:      false,
		param2:      false,
		param3:      false,
		pointer:     0,
		instruction: 0,
	}

	ic.ParseNextInstruction()

	return ic, nil
}

// ParseNextInstruction parses the instruction found at pointer location in program array
func (ic *IntCode) ParseNextInstruction() {
	pointerValue := ic.program[ic.pointer]
	ic.instruction = pointerValue % 100
	if pointerValue > 10000 {
		ic.param1 = true
		pointerValue -= 10000
		if pointerValue > 1000 {
			ic.param2 = true
			pointerValue -= 1000
			if pointerValue > 100 {
				ic.param3 = true
			}
		}
	}
	if pointerValue > 1000 {
		ic.param1 = true
		pointerValue -= 1000
		if pointerValue > 100 {
			ic.param2 = true
		}
	}
	if pointerValue > 100 {
		ic.param1 = true
	}
}

func parseIntArrayFromString(input string) ([]int, error) {
	var arr []int
	for _, num := range strings.Split(input, ",") {
		conv, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Failed to convert num: %s", num)
		}
		arr = append(arr, conv)
	}
	return arr, nil
}
