package main

import (
	"bufio"
	"fmt"
	"os"
)
 
type groupAnswers [26]bool

func scoreGroup(ga groupAnswers) (score int) {
	for _, answer := range ga {
		if answer {
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
				gAnswers[r - 'a'] = true
			}
		}
		total += scoreGroup(gAnswers)
	}

	fmt.Println(total)
}