package main

const (
	MOV = "mov" // Move source (register or value) to destination
	OUT = "out" // Print
	ADD = "add" // Addition
	SUB = "sub" // Substraction
	NEG = "neg" // Arithmetic negation
	NOT = "not" // Bitwise complement
	NOR = "nor"
	AND = "and"
	OR  = "or"
	XOR = "xor"
	INC = "inc" // Increment by 1
	DEC = "dec" // Decrement by 1
	SHF = "shl" // Left shift
	JMP = "jmp" // Jump to label
	JZ  = "jz"  // Jump if zero
	JNZ = "jnz" // Jump if not zero
)

const REG_SIZE = 64

type asm struct {
	registers [REG_SIZE]uint8
}
