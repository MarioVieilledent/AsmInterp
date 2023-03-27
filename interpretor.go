package main

import (
	"fmt"
	"strconv"
	"strings"
)

func interpret(fileContent string) {
	// Split the content into lines
	fileLines := strings.Split(fileContent, "\n")

	// Loop through each line and print its content
	for lineNumber, lineString := range fileLines {
		// fmt.Println("r0 = " + strconv.Itoa(int(a.registers[0])))

		// Split the line into words
		words := strings.Fields(strings.ToLower(lineString))

		// Use the first word as the switch value
		switch words[0] {
		case MOV:
			if len(words) == 3 {
				reg, isValidReg := isValidRegister(words[1])
				if !isValidReg {
					exit(lineNumber, MOV+" instruction, \""+words[1]+"\" is not a valid register.")
				}

				val, isValidVal := isValidValue(words[2])
				if !isValidVal {
					reg2, isValidReg2 := isValidRegister(words[2])
					if !isValidReg2 {
						exit(lineNumber, MOV+" instruction, \""+words[2]+"\" is not a valid value or register.")
					} else {
						a.registers[reg] = a.registers[reg2]
					}
				} else {
					a.registers[reg] = val
				}
			} else {
				exit(lineNumber, MOV+" instruction requires 2 arguments.")
			}
		case OUT:
			if len(words) == 2 {
				v, is := isValidValueOrRegister(words[1])
				if is == 1 {
					fmt.Println(v)
				} else if is == 2 {
					fmt.Println(a.registers[v])
				} else {
					exit(lineNumber, OUT+" instruction, \""+words[1]+"\" is not a valid value or register.")
				}
			} else {
				exit(lineNumber, OUT+" instruction should have 1 operands.")
			}
		case ADD:
			fmt.Println("Matched AND")
		case SUB:
			fmt.Println("Matched SUB")
		case OR:
			fmt.Println("Matched OR")
		case AND:
			if len(words) == 4 {
				v1, is1 := isValidValueOrRegister(words[2])
				if is == 1 {
					fmt.Println(v)
				} else if is == 2 {
					fmt.Println(a.registers[v])
				} else {
					exit(lineNumber, AND+" instruction, \""+words[1]+"\" is not a valid value or register.")
				}

				v2, is2 := isValidValueOrRegister(words[3])
				if is == 1 {
					fmt.Println(v)
				} else if is == 2 {
					fmt.Println(a.registers[v])
				} else {
					exit(lineNumber, AND+" instruction, \""+words[1]+"\" is not a valid value or register.")
				}

				
				reg, isReg := isValidRegister(words[1])
				if isReg {
					a.registers[reg] =
				} else {
					exit(lineNumber, AND+" instruction, \""+words[1]+"\" is not a valid register.")
				}
			} else {
				exit(lineNumber, AND+" instruction should have 3 operands.")
			}
		case XOR:
			fmt.Println("Matched XOR")
		case INC:
			fmt.Println("Matched INC")
		case DEC:
			fmt.Println("Matched DEC")
		case SHF:
			fmt.Println("Matched SHF")
		case JMP:
			fmt.Println("Matched JMP")
		case JZ:
			fmt.Println("Matched JIZ")
		default:
			exit(lineNumber, "Unknown instruction: \""+words[0]+"\".")
		}
	}
}

func isValidRegister(str string) (int, bool) {
	if len(str) < 2 || str[0] != 'r' {
		return 0, false
	}
	numStr := str[1:]
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 || num > 63 {
		return 0, false
	}
	return num, true
}

func isValidValue(str string) (uint8, bool) {
	if num, err := strconv.Atoi(str); err == nil && num >= 0 && num <= 255 {
		return uint8(num), true
	}
	if len(str) >= 3 && str[0:2] == "0x" {
		hexStr := str[2:]
		if num, err := strconv.ParseUint(hexStr, 16, 8); err == nil && num <= 255 {
			return uint8(num), true
		}
	}
	return 0, false
}

// 0 = error, 1 = valid value, 2 = valid register
func isValidValueOrRegister(str string) (uint8, int) {
	if str[0] != 'r' {
		if num, err := strconv.Atoi(str); err == nil && num >= 0 && num <= 255 {
			return uint8(num), 1
		}
		if len(str) >= 3 && str[0:2] == "0x" {
			hexStr := str[2:]
			if num, err := strconv.ParseUint(hexStr, 16, 8); err == nil && num <= 255 {
				return uint8(num), 1
			}
		}
		return 0, 0
	}
	numStr := str[1:]
	num, err := strconv.Atoi(numStr)
	if err != nil || num < 0 || num > 63 {
		return 0, 0
	}
	return uint8(num), 2
}

func exit(lineNumber int, message string) {
	fmt.Println("Error at line " + strconv.Itoa(lineNumber+1) + ": " + message)
}
