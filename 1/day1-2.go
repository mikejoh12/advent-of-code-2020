package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
    var nrs []int
  
    for scanner.Scan() {
		nrStr := scanner.Text()
		nr, _ := strconv.Atoi(nrStr)
        nrs = append(nrs, nr)
    }

	for i := 0; i < len(nrs); i++ {
		for j := i + 1; j < len(nrs); j++ {
			for k := j + 1; k < len(nrs); k++ {
				if nrs[i] + nrs[j] + nrs[k] == 2020 {
					fmt.Println(nrs[i] * nrs[j] * nrs[k])
				}
			}
		}
	}

	file.Close()
}