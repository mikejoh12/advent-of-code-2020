package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func validate(s string) bool {
	var fields []string = []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}

	for _, field := range fields {
		if !strings.Contains(s, field) {
			return false
		}
	}

	return true
}
 
func main() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
  
	var dataStr string

    for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			dataStr += "\n"
		} else {
			dataStr += s + " "
		}
    }
	dataStr = strings.TrimSpace(dataStr)

	var data []string = strings.Split(dataStr, "\n")
	var validCount int

	for _, d := range data {
		if validate(d) {
			validCount++
		}
	}

	fmt.Println("count", validCount)
	file.Close()
}