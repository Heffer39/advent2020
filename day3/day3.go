package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func readGeologyData() []string {
	abs, err := filepath.Abs("./localGeology.txt")
	file, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var data []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	if len(data) < 1 {
		log.Fatal("data not parsed")
	}
	return data
}

func tobogganSlope(data []string, right int, down int) int {
	var tree int
	var index int
	for i, line := range data {
		if i == 0 || i%down != 0 {
			continue
		}
		index += right
		if index >= len(line) {
			index = index - len(line)
		}
		if line[index] == '#' {
			tree++
		}
	}
	return tree
}
