package main

import (
	"fmt"
)

func twoSum(data []int) (int, error) {
	frontIndex := 0
	backIndex := len(data) - 1
	for {
		frontVal := data[frontIndex]
		backVal := data[backIndex]
		sum := frontVal + backVal
		if sum == 2020 {
			multiply := frontVal * backVal
			return multiply, nil
		}
		if sum > 2020 {
			backIndex--
		}
		if sum < 2020 {
			frontIndex++
		}
		if frontIndex >= backIndex {
			return 0, fmt.Errorf("no result found")
		}
	}
}
