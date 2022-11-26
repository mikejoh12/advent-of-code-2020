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
    var total int
  
    for scanner.Scan() {
		s := scanner.Text()
		s = strings.Replace(s, "-", " ", 1)
		s = strings.Replace(s, ":", "", 1)

		r := strings.NewReader(s)

		var min, max int
		var l, pwd string
		_, err := fmt.Fscanf(r, "%d %d %s %s", &min, &max, &l, &pwd)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
		}

		count := strings.Count(pwd, l)
		if count >= min && count <= max {
			total++
		}

    }
	fmt.Println("total", total)
	file.Close()
}