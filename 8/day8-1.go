package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	operation	string
	argument	int
	isExecuted	bool
}

func main() {
	file, _ := os.Open("input.txt")
    scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	
	var code []instruction

    for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " ")
		arg, _ := strconv.Atoi(parts[1])
		code = append(code, instruction{operation: parts[0], argument: arg})
	}

	i, acc := 0, 0
	for {
		instr := code[i]

		if instr.isExecuted {
			fmt.Println("acc", acc)
			break
		}
		code[i].isExecuted = true

		switch instr.operation {
		case "nop":
			i++
		case "jmp":
			i += instr.argument
		case "acc":
			acc += instr.argument
			i++
		}
	}
}