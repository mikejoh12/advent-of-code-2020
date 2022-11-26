package main

import (
	"bufio"
	"fmt"
	"os"
)
 
type groupAnswers [26]int

func scoreGroup(ga groupAnswers, gSize int) (score int) {
	for _, answer := range ga {
		if answer == gSize {
			score++
		}
	}
	return
}

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var data [][]string = [][]string{{}}
	var total int

    for scanner.Scan() {
		s := scanner.Text()
		if s != "" {
			data[len(data)-1] = append(data[len(data)-1], s)
		} else {
			scanner.Scan()
			s = scanner.Text()
			data = append(data, []string{s})
		}
	}
	
	for _, groupData := range data {
		var gAnswers groupAnswers
		for _, personData := range groupData {
			for _, r := range personData {
				gAnswers[r - 'a']++
			}
		}
		total += scoreGroup(gAnswers, len(groupData))
	}

	fmt.Println(total)
}