package main

const (
	MOV = "mov" // Move source (register or value) to destination
	OUT = "out" // Print
	ADD = "add" // Addition
	SUB = "sub" // Subtraction
	NEG = "neg" // Arithmetic negation
	NOT = "not" // Bitwise complement
	AND = "and" // Bitwise AND
	OR  = "or"  // Bitwise OR
	XOR = "xor" // Bitwise XOR
	INC = "inc" // Increment
	DEC = "dec" // Decrement by 1
	SHL = "shl" // Left shift
	SHR = "shr" // Right shift
	JMP = "jmp" // Jump
	JZ  = "jz"  // Jump if zero
	JNZ = "jnz" // Jump if not zero
)

const REG_SIZE = 64

type asm struct {
	reg [REG_SIZE]uint8 // Registers
	adr uint8           // Pointer to instruction
}
