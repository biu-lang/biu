# runtime
.text
.globl _start
_start:
	movq $1, %rdi
	movq $60, %rax	# exit
	syscall

