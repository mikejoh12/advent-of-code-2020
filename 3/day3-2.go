package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func countMap(m [][]string, xDiff, yDiff int) int {
	var y, x, count int

	for y < len(m) {
		if m[y][x] == "#" {
			count++
		}
		y, x = y + yDiff, (x + xDiff) % len(m[0]) 
	}

	return count
}

func main() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
  
	var partMap [][]string

    for scanner.Scan() {
		s := scanner.Text()
		partMap = append(partMap, strings.Split(s, ""))
    }

	c := countMap(partMap, 1, 1)
	c *= countMap(partMap, 3, 1)
	c *= countMap(partMap, 5, 1)
	c *= countMap(partMap, 7, 1)
	c *= countMap(partMap, 1, 2)

	fmt.Println("count", c)
	file.Close()
}