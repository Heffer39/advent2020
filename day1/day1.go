package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	abs, err := filepath.Abs("./sumdata.txt")
	file, err := os.Open(abs)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, _ := strconv.Atoi(scanner.Text())
		data = append(data, i)
	}

	if len(data) < 1 {
		log.Fatal("data not parsed")
	}
	sort.Ints(data)

	twoSum, err := twoSum(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("twoSum: %v\n", twoSum)

	threeSum, err := threeSum(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("threeSum: %v\n", threeSum)
}
