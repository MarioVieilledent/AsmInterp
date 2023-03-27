package main

const (
	MOV = "mov" // Move source (register or value) to destination
	OUT = "out" // Print
	ADD = "add" // Addition
	SUB = "sub" // Subtraction
	NEG = "neg" // Arithmetic negation
	NOT = "not" // Bitwise complement
	AND = "and" // Logic gate AND
	OR  = "or"  // Logic gate OR
	XOR = "xor" // Logic gate XOR
	INC = "inc" // Increment by 1
	DEC = "dec" // Decrement by 1
	SHF = "shl" // Left shift
	JMP = "jmp" // Jump to label
	JZ  = "jz"  // Jump if zero
	JNZ = "jnz" // Jump if not zero
)

const REG_SIZE = 64

type asm struct {
	reg [REG_SIZE]uint8 // Registers
	adr uint8           // Pointer to instruction
}
