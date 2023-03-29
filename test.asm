mov r0 245
mov r1 10
out r1
shl r1 1
out r1
inc r0
jnz 3

mov r0 245
mov r1 255
out r1
shr r1 1
out r1
inc r0
jnz 11