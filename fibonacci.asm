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