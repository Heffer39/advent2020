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
		if sign == "-" {
			value *= -1
		}
		instructionsList = append(instructionsList, instruction{
			operation: operation,
			value:     value,
			index:     i,
		})
	}
	found, accumulator := runOperations(
		instructionsList, 0, 0, false)

	fmt.Printf("accumulator: %v, found: %v", accumulator, found)
}

func runOperations(instructionsList []instruction, accumulator, index int, inception bool) (found bool, result int) {
	instructionsMap := make(map[int]*instruction)

	for {
		if index >= len(instructionsList) {
			fmt.Printf("reached end of list!\n")
			return true, accumulator
		}
		if _, ok := instructionsMap[index]; ok {
			return false, accumulator
		}
		ins := instructionsList[index]
		instructionsMap[index] = &ins
		switch ins.operation {
		case "nop":
			if !inception {
				loop, _ := runOperations(instructionsList, accumulator, index+ins.value, true)
				if loop {
					fmt.Printf("NOP! broken index: %v, operation: %v, value: %v\n", ins.index+1, ins.operation, ins.value)
					return true, accumulator
				}
			}
			index++
		case "acc":
			accumulator += ins.value
			index++
		case "jmp":
			if !inception {
				loop, _ := runOperations(instructionsList, accumulator, index+1, true)
				if loop {
					fmt.Printf("JUMP! broken index: %v, operation: %v, value: %v\n", ins.index+1, ins.operation, ins.value)
					return true, accumulator
				}
			}
			index += ins.value
		}
	}
}

type instruction struct {
	operation    string
	index, value int
}
