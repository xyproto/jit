.PHONY: all run clean

all: hello.mc run

hello.o: hello.asm
	nasm -f elf64 -o hello.o hello.asm

hello.bin: hello.o
	objcopy -O binary --only-section=.text hello.o hello.bin

hello.mc: hello.bin
	hexdump -C hello.bin | cut -b11- | cut -d'|' -f1 > hello.mc

run:
	./jitrun hello.mc

clean:
	rm hello.bin hello.o hello.mc
