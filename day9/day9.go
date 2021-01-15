package main

import (
	"advent2020/lib"
	"fmt"
	"strconv"
)

// main function design:
// Two pointers to denote beginning and end of 25 number preamble
// Map to store numbers in the current preamble map[int]int
// key is the number being stored, value is the count of how many we've seen - to track duplicates
// as we move the pointers, decrement the old key and add the new one
// if old key value == 0, delete
// for each new value in the leading pointer, iterate backwards, running map[sum - index]
// if it exists, we've found a match and can move on
// if no match exists by the time we check the back pointer, we've found out invalid data
func main() {
	data := lib.ReadFile("XMAS.txt")

	var xmas []int
	for _, s := range data {
		v, _ := strconv.Atoi(s)
		xmas = append(xmas, v)
	}

	invalidNumber := findInvalidNumber(xmas)
	fmt.Printf("invalid index: %v, value: %v\n", invalidNumber.index+1, xmas[invalidNumber.index])
	encryptionWeakness := findContiguousSum(xmas, invalidNumber)
	fmt.Printf("encyrption weakness: %v", encryptionWeakness)
}

func findInvalidNumber(xmas []int) (invalidNumber number) {
	preambleSize := 25
	xmasMap := make(map[int]int)
	valid := false
	for index := 0; index < len(xmas); index++ {
		if index >= preambleSize {
			back := index - preambleSize
			for ; back < index; back++ {
				targetVal := xmas[index] - xmas[back]
				if _, ok := xmasMap[targetVal]; ok {
					valid = true
					break
				}
			}
			if !valid {
				invalidNumber = number{
					index: index,
					value: xmas[index],
				}
				break
			}
			valid = false

			xmasMap[back]--
			if xmasMap[back] <= 0 {
				delete(xmasMap, back)
			}
		}
		xmasMap[xmas[index]]++
	}
	return
}

func findContiguousSum(xmas []int, invalidNumber number) int {
	var weakness pair
	for index := 0; index < invalidNumber.index; index++ {
		sum := xmas[index]
		for frontIndex := index + 1; frontIndex < invalidNumber.index; frontIndex++ {
			sum += xmas[frontIndex]
			if sum == invalidNumber.value {
				weakness = pair{
					frontIndex: frontIndex,
					backIndex:  index,
				}
				break
			}
			if sum > invalidNumber.value {
				break
			}
		}
	}

	var smallest, largest int
	for index := weakness.backIndex; index <= weakness.frontIndex; index++ {
		val := xmas[index]
		if val > largest {
			largest = val
		}
		if val < smallest || smallest == 0 {
			smallest = val
		}
	}
	result := smallest + largest
	return result
}

type pair struct {
	frontIndex, backIndex int
}

type number struct {
	index, value int
}
