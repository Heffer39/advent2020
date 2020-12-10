package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	abs, err := filepath.Abs("./passwordData.txt")
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

	oldSledPasswords := oldSledRentalPasswordPolicy(data)
	fmt.Printf("old sled valid passwords: %v\n", oldSledPasswords)

	tobogganPasswords := tobogganPasswordPolicy(data)
	fmt.Printf("toboggan valid passwords: %v", tobogganPasswords)
}

func tobogganPasswordPolicy(data []string) int {
	validPasswordCount := 0
	for _, line := range data {
		policy := strings.Split(line, " ")
		minMax := strings.Split(policy[0], "-")
		firstPosition, _ := strconv.Atoi(minMax[0])
		secondPosition, _ := strconv.Atoi(minMax[1])
		policyLetter := rune(policy[1][0])
		password := policy[2]
		policyCount := 0

		if policyLetter == rune(password[firstPosition-1]) {
			policyCount++
		}
		if policyLetter == rune(password[secondPosition-1]) {
			policyCount++
		}
		if policyCount == 1 {
			validPasswordCount++
		}
	}
	return validPasswordCount
}

func oldSledRentalPasswordPolicy(data []string) int {
	validPasswordCount := 0
	for _, line := range data {
		policy := strings.Split(line, " ")
		minMax := strings.Split(policy[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		policyLetter := rune(policy[1][0])
		password := policy[2]
		policyCount := 0
		for _, c := range password {
			if c == policyLetter {
				policyCount++
			}
		}
		if policyCount <= max && policyCount >= min {
			validPasswordCount++
		}
	}
	return validPasswordCount
}
