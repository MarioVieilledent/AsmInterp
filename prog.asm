mov r0 246

mov r1 0x1
mov r2 1
out r1
out r2

add r3 r1 r2
mov r1 r2
mov r2 r3
out r3
inc r0
jnz 7