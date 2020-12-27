package main

import (
	"advent2020/lib"
	"fmt"
	"log"
	"unicode"
)

func main() {
	data := lib.ReadFile("questions.txt")

	questions := make(map[rune]int)
	var anyoneYesCount []int
	var everyoneYesCount []int
	var groupSize int

	for i, s := range data {
		groupSize++
		for _, c := range s {
			if !unicode.IsLetter(c) {
				log.Fatal("incorrect input")
			}
			questions[c]++
		}
		if s == "" || i == len(data)-1 {
			anyoneCount := 0
			everyoneCount := 0
			for _, v := range questions {
				anyoneCount++
				if v == groupSize-1 || (v == groupSize && i == len(data)-1) {
					everyoneCount++
				}
			}
			if anyoneCount != 0 {
				anyoneYesCount = append(anyoneYesCount, anyoneCount)
			}
			if everyoneCount != 0 {
				everyoneYesCount = append(everyoneYesCount, everyoneCount)
			}
			questions = make(map[rune]int)
			groupSize = 0
			continue
		}
	}
	var anyoneSum, everyoneSum int
	for _, v := range anyoneYesCount {
		anyoneSum += v
	}
	for _, v := range everyoneYesCount {
		everyoneSum += v
	}
	fmt.Printf("anyone: %v, everyone: %v\n", anyoneSum, everyoneSum)
}
