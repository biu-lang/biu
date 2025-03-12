# Runtime
.text
.globl _start
_start:
	call main.main

# .globl os.Exit
os.Exit:
	movq 8(%rsp), %rdi	# arg1
	movq $60, %rax			# exit
	syscall

# End of program

