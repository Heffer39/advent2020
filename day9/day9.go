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
				fmt.Printf("invalid index: %v, value: %v", index+1, xmas[index])
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
}
