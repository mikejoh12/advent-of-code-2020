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

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  
	var maxSeatId int

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
		if rows[0] * 8 + cols[0] > maxSeatId {
			maxSeatId = rows[0] * 8 + cols[0]
		}
    }
	
	fmt.Println(maxSeatId)
}