# Run this on Linux
#

AS ?= as
LD ?= ld

biu: runtime.o
	$(LD) -o $@ $^

runtime.o: runtime.s
	$(AS) -o $@ $^

test: biu
	./test.sh

.PHONY: test clean
clean:
	$(RM) -f *.o biu

