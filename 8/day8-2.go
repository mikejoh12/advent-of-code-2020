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

func runCode(c []instruction) (int, bool) {
	i, acc := 0, 0
	for {
		instr := c[i]

		if instr.isExecuted {
			return -1, false
		}
		c[i].isExecuted = true

		switch instr.operation {
		case "nop":
			i++
		case "jmp":
			i += instr.argument
		case "acc":
			acc += instr.argument
			i++
		}

		if i == len(c) {
			return acc, true
		}
	}
}

func modifyCode(c []instruction, line int) []instruction {
	newCode := make([]instruction, len(c))
	copy(newCode, c)
	if newCode[line].operation == "jmp" {
		newCode[line].operation = "nop"
	} else if newCode[line].operation == "nop" {
		newCode[line].operation = "jmp"
	}
	return newCode
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

	for i, instr := range code {
		if instr.operation == "jmp" || instr.operation == "nop" {
			newCode := modifyCode(code, i)
			if nr, ok := runCode(newCode); ok {
				fmt.Println(nr)
				break
			}
		}
	}
}