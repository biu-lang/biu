# Run this on Linux
#

AS ?= as
LD ?= ld

biu: main.o
	$(LD) -o $@ $^

main.s: main.go
	go run main.go > main.s

%.o:%.s
	$(AS) -o $@ $^

main.o: main.s runtime.s


test: biu
	./test.sh

.PHONY: test clean
clean:
	$(RM) -f *.o biu main.s

