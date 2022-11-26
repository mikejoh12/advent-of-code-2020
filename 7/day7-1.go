package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bagParents map[string][]string

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bMap := make(bagParents)

    for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(strings.Trim(s, "."), " bags contain ")
		children := strings.Split(parts[1], ", ")
		parent := parts[0]

		for _, child := range children {
			if child != "no other bags" {
				tChild := strings.Replace(child, " bags", "", -1)
				tChild = strings.Replace(tChild, " bag", "", -1)
				bMap[tChild[2:]] = append(bMap[tChild[2:]], parent)
			}
		}
	}

	var find func(bag string)
	parentMap := make(map[string]bool)

	find = func(bag string) {
		if parents, ok := bMap[bag]; ok {
			for _, parent := range parents {
				parentMap[parent] = true
				find(parent)
			}
		}
	}
	
	find("shiny gold")
	fmt.Println(len(parentMap))
}