package main

import (
	"fmt"
)

func threeSum(data []int) (int, error) {
	frontIndex := 0
	for {
		frontVal := data[frontIndex]
		for midIndex, backIndex := frontIndex + 1, len(data) - 1; midIndex < backIndex; {
			midVal := data[midIndex]
			backVal := data[backIndex]
			sum := frontVal + midVal + backVal
			if sum == 2020 {
				multiply := frontVal * backVal * midVal
				return multiply, nil
			}
			if sum > 2020 {
				backIndex--
			}
			if sum < 2020 {
				midIndex++
			}
		}
		if frontIndex >= len(data) - 1 {
			return 0, fmt.Errorf("no result found")
		}
		frontIndex++
	}
}
