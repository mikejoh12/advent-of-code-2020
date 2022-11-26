package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bagData struct {
	name	string
	qty		int
}	

type bagChildren map[string][]bagData

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	bMap := make(bagChildren)

    for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(strings.Trim(s, "."), " bags contain ")
		children := strings.Split(parts[1], ", ")
		parent := parts[0]

		for _, child := range children {
			if child != "no other bags" {
				tChild := strings.Replace(child, " bags", "", -1)
				tChild = strings.Replace(tChild, " bag", "", -1)

				name := tChild[2:]
				qty, _ := strconv.Atoi(tChild[0:1])

				bMap[parent] = append(bMap[parent], bagData{name: name, qty: qty})
			}
		}
	}

	var find func(bag string)
	var bagCount int

	find = func(bag string) {
		if children, ok := bMap[bag]; ok {
			for _, child := range children {
				for i := 1; i <= child.qty; i++ {
					bagCount++
					find(child.name)
				}
			}
		}
	}
	
	find("shiny gold")
	fmt.Println(bagCount)
}