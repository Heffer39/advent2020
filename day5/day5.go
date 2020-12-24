package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type seat struct {
	row, col, seatID int
}

func main() {
	abs, err := filepath.Abs("./boardingPasses.txt")
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

	var highestBoardingPass, lowestBoardingPass int
	var minRow, maxRow int
	var minCol, maxCol int
	var seatMap = make(map[int]seat)

	for _, s := range data {
		if len(s) != 10 {
			log.Printf("incorrect size: %v\n", s)
			continue
		}
		minRow, maxRow, minCol, maxCol = 0, 127, 0, 7

		for _, c := range s {
			rowMedian := (maxRow + minRow) / 2
			colMedian := (maxCol + minCol) / 2

			switch c {
			case 'F':
				maxRow = rowMedian
			case 'B':
				minRow = rowMedian + 1
			case 'L':
				maxCol = colMedian
			case 'R':
				minCol = colMedian + 1
			}
		}

		seatID := minRow*8 + minCol
		//fmt.Printf("row: %v, col: %v, seatID: %v\n", minRow, minCol, seatID)
		if seatID > highestBoardingPass {
			highestBoardingPass = seatID
		}
		if seatID < lowestBoardingPass || lowestBoardingPass == 0 {
			lowestBoardingPass = seatID
		}
		seatMap[seatID] = seat{row: minRow, col: minCol, seatID: seatID}
	}

	for i := lowestBoardingPass; i <= highestBoardingPass; i++ {
		if _, ok := seatMap[i]; !ok {
			fmt.Printf("missing val: %v\n", i)
		}
	}

	fmt.Printf("highest boarding pass: %v, lowest: %v", highestBoardingPass, lowestBoardingPass)
}
