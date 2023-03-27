package main

import (
	"fmt"
	"strconv"
	"strings"
)

func interpLines(fileContent string) {
	// Split the content into lines
	fileLines := strings.Split(fileContent, "\n")

	for int(a.adr) < len(fileLines) {
		if strings.TrimSpace(fileLines[a.adr]) != "" {
			interp(fileLines[a.adr], int(a.adr))
		}
		a.adr++
	}
}

func interp(line string, ln int) {
	// Split the line into words
	words := strings.Fields(strings.ToLower(line))

	// Use the first word as the switch value
	switch words[0] {
	case MOV:
		if len(words) != 3 {
			exit(ln, MOV+" instruction requires 2 arguments.")
		}

		reg, isValidReg := isValidRegister(words[1])
		if !isValidReg {
			exit(ln, MOV+" instruction, \""+words[1]+"\" is not a valid register.")
		}

		val, isValidVal := isValidValue(words[2])
		if !isValidVal {
			reg2, isValidReg2 := isValidRegister(words[2])
			if !isValidReg2 {
				exit(ln, MOV+" instruction, \""+words[2]+"\" is not a valid value or register.")
			} else {
				a.reg[reg] = a.reg[reg2]
			}
		} else {
			a.reg[reg] = val
		}
	case OUT:
		if len(words) == 2 {
			reg, isValidReg := isValidRegister(words[1])
			if isValidReg {
				fmt.Println(a.reg[reg])
			} else {
				exit(ln, OUT+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, OUT+" instruction should have 1 operand.")
		}
	case ADD:
		if len(words) == 4 {
			reg1, isValidReg1 := isValidRegister(words[1])
			if isValidReg1 {
				reg2, isValidReg2 := isValidRegister(words[2])
				if isValidReg2 {
					reg3, isValidReg3 := isValidRegister(words[3])
					if isValidReg3 {
						a.reg[reg1] = a.reg[reg2] + a.reg[reg3]
					} else {
						exit(ln, ADD+" instruction, \""+words[3]+"\" is not a valid value or register.")
					}
				} else {
					exit(ln, ADD+" instruction, \""+words[2]+"\" is not a valid value or register.")
				}
			} else {
				exit(ln, ADD+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, ADD+" instruction should have 1 operand.")
		}
	case SUB:
		if len(words) == 4 {
			reg1, isValidReg1 := isValidRegister(words[1])
			if isValidReg1 {
				reg2, isValidReg2 := isValidRegister(words[2])
				if isValidReg2 {
					reg3, isValidReg3 := isValidRegister(words[3])
					if isValidReg3 {
						a.reg[reg1] = a.reg[reg2] - a.reg[reg3]
					} else {
						exit(ln, SUB+" instruction, \""+words[3]+"\" is not a valid value or register.")
					}
				} else {
					exit(ln, SUB+" instruction, \""+words[2]+"\" is not a valid value or register.")
				}
			} else {
				exit(ln, SUB+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, SUB+" instruction should have 1 operand.")
		}
	case OR:
		fmt.Println("Matched OR")
	case AND:
		if len(words) != 4 {
		} else {
			exit(ln, AND+" instruction should have 3 operands.")
		}
	case XOR:
		fmt.Println("Matched XOR")
	case INC:
		if len(words) == 2 {
			reg, isValidReg := isValidRegister(words[1])
			if isValidReg {
				a.reg[reg]++
			} else {
				exit(ln, INC+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, INC+" instruction should have 1 operand.")
		}
	case DEC:
		if len(words) == 2 {
			reg, isValidReg := isValidRegister(words[1])
			if isValidReg {
				a.reg[reg]--
			} else {
				exit(ln, INC+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, INC+" instruction should have 1 operand.")
		}
	case SHF:
		fmt.Println("Matched SHF")
	case JMP:
		if len(words) == 2 {
			reg, isValidReg := isValidRegAddress(words[1])
			if isValidReg {
				a.adr = uint8(reg) - 1
			} else {
				exit(ln, JMP+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, JMP+" instruction should have 1 operand.")
		}
	case JZ:
		if len(words) == 2 {
			reg, isValidReg := isValidRegAddress(words[1])
			if isValidReg {
				if a.reg[0] == 0 {
					a.adr = uint8(reg) - 1
				}
			} else {
				exit(ln, JZ+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, JZ+" instruction should have 1 operand.")
		}
	case JNZ:
		if len(words) == 2 {
			reg, isValidReg := isValidRegAddress(words[1])
			if isValidReg {
				if a.reg[0] != 0 {
					a.adr = uint8(reg) - 1
				}
			} else {
				exit(ln, JNZ+" instruction, \""+words[1]+"\" is not a valid value or register.")
			}
		} else {
			exit(ln, JNZ+" instruction should have 1 operand.")
		}
	default:
		exit(ln, "Unknown instruction: \""+words[0]+"\".")
	}
}

func isValidRegAddress(str string) (int, bool) {
	num, err := strconv.Atoi(str)
	if err != nil || num < 0 || num > 63 {
		return 0, false
	}
	return num, true
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

func exit(ln int, message string) {
	fmt.Println("Error at line " + strconv.Itoa(ln) + ": " + message)
}
