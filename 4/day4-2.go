package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isValidYear(s string, min, max int) bool {
	yrRe := regexp.MustCompile(`\d{4}`)
	yr, _ := strconv.Atoi(yrRe.FindString(s))
	if yr >= min && yr <= max {
		return true
	}
	return false
}

func valHeight(s string) bool {
	s = strings.TrimLeft(s, "hgt:")
	heightRe := regexp.MustCompile(`\d*`)
	heightStr := heightRe.FindString(s)
	height, _ := strconv.Atoi(heightRe.FindString(heightStr))
	unitRe := regexp.MustCompile(`[a-zA-Z]{2}`)
	unit := unitRe.FindString(s)

	if unit == "cm" && height >= 150 && height <= 193 {
		return true
	}

	if unit == "in" && height >= 59 && height <= 76 {
		return true
	}
	return false
}

func validate(s string) bool {

	byrRe := regexp.MustCompile(`byr:\d\d\d\d`)
	iyrRe := regexp.MustCompile(`iyr:\d\d\d\d`)
	eyrRe := regexp.MustCompile(`eyr:\d\d\d\d`)
	hgtRe := regexp.MustCompile(`hgt:\d*(in|cm)`)
	hclRe := regexp.MustCompile(`hcl:#[0-9a-f]{6}`)
	eclRe := regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)`)
	pidRe := regexp.MustCompile(`pid:\d{9}`)

	tooLongPidRe := regexp.MustCompile(`pid:\d{10}`)
	tooLongPid := tooLongPidRe.FindString(s)

	byr := byrRe.FindString(s)
	iyr := iyrRe.FindString(s)
	eyr := eyrRe.FindString(s)
	hgt := hgtRe.FindString(s)
	hcl := hclRe.FindString(s)
	ecl := eclRe.FindString(s)
	pid := pidRe.FindString(s)

	if byr == "" || iyr == "" || eyr == "" || hgt == "" || hcl == "" || ecl == "" || pid == "" {
		return false
	}

	if tooLongPid != "" {
		return false
	}

	if !isValidYear(byr, 1920, 2002) || !isValidYear(iyr, 2010, 2020) || !isValidYear(eyr, 2020, 2030) {
		return false
	}

	if !valHeight(hgt) {
		return false
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

	fmt.Println("valid count", validCount)
	file.Close()
}