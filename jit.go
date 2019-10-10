package jit

// #include <signal.h>
// #include <stdio.h>
// #include <assert.h>
// #include <stdlib.h>
// #include <stdbool.h>
// #include <stdint.h>
// #include <stdbool.h>
// #include <errno.h>
//
// static int illegal = 0;
// static int segfault = 0;
//
// void sighandler(int sig) {
//   switch (sig) {
//   case SIGILL:
//     illegal = 1;
//     fputs("Caught SIGILL: illegal instruction <3\n", stderr);
//     break;
//   case SIGSEGV:
//     segfault = 1;
//     fputs("Caught SIGSEGV\n", stderr);
//     break;
//   }
// }
//
// int run(void* mem) {
//   illegal = 0;
//   segfault = 0;
//   signal(SIGILL, sighandler);
//   signal(SIGSEGV, sighandler);
//   int (*f)() = mem;
//   return f();
// }
//
// int getillegal() {
//   return illegal;
// }
//
// int getsegfault() {
//   return segfault;
// }
import "C"
import (
	"errors"
	"syscall"
	"unsafe"
)

// ErrMemCopy may be returned if code could not be copied over to "mmapped" memory
var ErrMemCopy = errors.New("could not copy code over to the executable memory area")

// Execute machinecode and return the value returned by the machinecode.
func Execute(code []byte) (int, error) {
	// Create a block of memory that can be executed if it is filled with machinecode
	executableArea, err := syscall.Mmap(0, 0, len(code), syscall.PROT_WRITE|syscall.PROT_EXEC, syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	if err != nil {
		return 0, err
	}

	// Copy the code over to the executable area
	bytesWritten := copy(executableArea, code)
	if bytesWritten != len(code) {
		return 0, ErrMemCopy
	}

	// Executing the code in memory directly from Go did not work out (it segfaults)
	//p := unsafe.Pointer(&executableArea[0])
	//f := (*func() int)(p)
	//return (*f)(), nil

	// Execute the code in memory, by using the C snippet at the top
	return int(C.run(unsafe.Pointer(&executableArea[0]))), nil
}

func GotSIGILL() bool {
	return C.getillegal() == 1
}

func GotSIGSEGV() bool {
	return C.getsegfault() == 1
}
