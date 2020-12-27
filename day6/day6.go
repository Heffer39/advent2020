package main

import (
	"advent2020/lib"
	"fmt"
	"log"
	"unicode"
)

func main() {
	data := lib.ReadFile("questions.txt")

	questions := make(map[rune]bool)
	var questionCounts []int

	for i, s := range data {
		for _, c := range s {
			if !unicode.IsLetter(c) {
				log.Fatal("incorrect input")
			}
			questions[c] = true
		}
		if s == "" || i == len(data)-1 {
			count := 0
			for range questions {
				count++
			}
			questionCounts = append(questionCounts, count)
			questions = make(map[rune]bool)
			continue
		}
	}
	var sum int
	for _, v := range questionCounts {
		sum += v
	}
	fmt.Println(sum)
}
