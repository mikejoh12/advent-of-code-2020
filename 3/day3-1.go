package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
  
	var partMap [][]string

    for scanner.Scan() {
		s := scanner.Text()
		partMap = append(partMap, strings.Split(s, ""))
    }

	var x, y, count int

	for y < len(partMap) {
		if partMap[y][x] == "#" {
			count++
		}
		y, x = y + 1, (x + 3) % len(partMap[0]) 
	}

	fmt.Println("count", count)
	file.Close()
}