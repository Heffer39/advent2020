package main

import (
	"fmt"
	"testing"
)

func TestTobogganSlope(t *testing.T) {
	var tests = []struct {
		right, down int
		want        int
	}{
		{1, 1, 88},
		{3, 1, 145},
		{5, 1, 71},
		{7, 1, 90},
		{1, 2, 42},
	}

	data := readGeologyData()

	sum := 1
	for _, tt := range tests {
		testname := fmt.Sprintf("right:%d,down:%d", tt.right, tt.down)
		t.Run(testname, func(t *testing.T) {
			got := tobogganSlope(data, tt.right, tt.down)
			fmt.Printf("got: %v\n", got)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
			sum *= got
		})
	}
	fmt.Printf("sum: %v\n", sum)
}
