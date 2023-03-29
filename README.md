# AsmInterp

An assembly code interpreter in Go for playing with chatGPT.

## Run

1. Code an assembly-like program with the instructions described below in a file ending in .asm
2. Run the go command `go run . <filename.asm>`

## How does it work

- There's 8 registers of 8 bits each: r0 to r7
- Values can be expressed with a number from 0 to 255, or a two digit hex value, from 0x00 to 0xff
- Commentaries comes after `//` and are ignored at interpretation

## Instruction

| Mnemonic | Example     | Description                                                       |
| -------- | ----------- | ----------------------------------------------------------------- |
| **mov**  | `mov r1 42` | Put value 42 into r1                                              |
| **mov**  | `mov r1 r2` | Copy value of r2 and put it into r1                               |
| **out**  | `out r3`    | Print value of r3                                                 |
| **add**  | `add r1 r2` | Add source to destination                                         |
| **sub**  | `sub r1 r2` | Subtract source to destination                                    |
| **neg**  | `neg r4`    | Arithmetic negation                                               |
| **not**  | `not r4`    | Bitwise negation                                                  |
| **and**  | `and r4 r5` | Bitwise AND destination by source                                 |
| **or**   | `or r4 r5`  | Not coded yet                                                     |
| **xor**  | `xor r4 r5` | Not coded yet                                                     |
| **inc**  | `inc r0`    | Increment r0 by 1                                                 |
| **dec**  | `dec r0`    | Decrement r0 by 1                                                 |
| **shl**  | `shl r1 k`  | Left shift destination by k bits                                  |
| **shr**  | `shl r1 k`  | Right shift destination by k bits                                 |
| **jmp**  | `jmp 4`     | Jump to instruction number 4 (line code 5)                        |
| **jz**   | `jz 4`      | jump to instruction number 4 (line code 5) if r0 is equal to zero |
| **jnz**  | `jzn 4`     | jump to instruction number 4 (line code 5) if r0 not zero         |

## Example of code: Fibonacci

```asm
// Fibonacci sequence

mov r0 245 // Used for conditional repetition

mov r1 0x1
mov r2 1
out r1
out r2

add r3 r1 r2
mov r1 r2
mov r2 r3
out r3
inc r0 // Increment r0 by 1
jnz 9 // Jump to line 10 (instruction 9) if r0 != 0
```