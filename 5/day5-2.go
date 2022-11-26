package main

import (
	"bufio"
	"fmt"
	"os"
)
 
func makeRange() (r []int) {
	for i := 0; i < 128; i++ {
		r = append(r, i)
	}
	return r
} 

func makePlaneSeats() [][8]bool {
	s := make([][8]bool, 0)
	for r := 0; r < 128; r++ {
		s = append(s, [8]bool{false, false, false, false, false, false, false, false})
	}
	return s
}

func findSeatId(s [][8]bool) int {
	for r := 1; r < len(s)-1; r++ {
		for c := 0; c < 8; c++ {
			seatStatus := s[r][c]
			
			if seatStatus == false {
				lowerId := r * 8 + c - 1
				lowerRow := lowerId / 8
				lowerCol := lowerId % 8

				higherId := r * 8 + c + 1
				higherRow := higherId / 8
				higherCol := higherId % 8

				if s[lowerRow][lowerCol] == true && s[higherRow][higherCol] == true {
					return r * 8 + c
				}
			}
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  
	sMap := makePlaneSeats()

    for scanner.Scan() {
		s := scanner.Text()
		row := s[0:7]
		col := s[7:10]

		rows := makeRange()
		var cols = []int{0,1,2,3,4,5,6,7}

		for _, l := range row {
			if l == 'F' {
				rows = rows[:len(rows)/2]
			} else if l == 'B' {
				rows = rows[len(rows)/2:]
			}
		}

		for _, l := range col {
			if l == 'L' {
				cols = cols[:len(cols)/2]
			} else if l == 'R' {
				cols = cols[len(cols)/2:]
			}
		}

		sMap[rows[0]][cols[0]] = true
    }

	a := findSeatId(sMap)
	fmt.Println(a)
	
}