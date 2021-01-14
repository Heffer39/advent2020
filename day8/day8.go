package main

import (
	"advent2020/lib"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := lib.ReadFile("instructions.txt")

	var instructionsList []instruction

	for i, v := range data {
		split := strings.Split(v, " ")
		operation := split[0]
		sign := split[1][0:1]
		value, _ := strconv.Atoi(split[1][1:])
		//fmt.Printf("sign: %v\n", sign)
		if sign == "-" {
			value *= -1
		}
		instructionsList = append(instructionsList, instruction{
			operation: operation,
			value:     value,
			index:     i,
		})
	}

	/*for _, v := range instructionsList {
		fmt.Printf("instruction: %v\n", v)
	}*/

	instructionsMap := make(map[int]*instruction)

	accumulator := 0
	index := 0
	for {
		ins := instructionsList[index]
		if _, ok := instructionsMap[index]; ok {
			break
		}
		instructionsMap[index] = &ins
		//fmt.Printf("index: %v, operation: %v, value: %v\n", ins.index+1, ins.operation, ins.value)
		switch ins.operation {
		case "nop":
			index++
		case "acc":
			accumulator += ins.value
			index++
		case "jmp":
			index += ins.value
		}
		//fmt.Printf("\taccumulator: %v\n", accumulator)
	}
	fmt.Printf("accumulator: %v", accumulator)
}

type instruction struct {
	operation    string
	index, value int
}
