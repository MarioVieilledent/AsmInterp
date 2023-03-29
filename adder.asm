// Addition without using ADD instruction
// It works by making a XOR and AND of the two numbers
// then left shift AND result by 1
// And repeat the process

mov r0 8

mov r1 42 // A = 42
mov r2 69 // B = 69

mov r3 r1
mov r4 r1
xor r3 r2 // C = 42 XOR 69
and r4 r2 // D = 42 AND 69
shl r4 1  // D = D << 1
mov r1 r3
mov r2 r4
dec r0

jnz 10

out r1 // Print result