package jit

// int run(void* mem) {
//   int (*f)() = mem;
//   return f();
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
