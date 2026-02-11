//go:build amd64 && linux

package jit

/*
#include <stdint.h>

// uint64_t because all our types in RegisterType fit in a register
uint64_t trampoline(void* fptr, uint64_t arg) {
    int (*f)(int) = fptr;
    return f(arg);
}
*/
import "C"
import (
	"errors"
	"fmt"
	"math"
	"syscall"
	"unsafe"
)

// ErrMemCopy may be returned if code could not be copied over to "mmapped" memory
var ErrMemCopy = errors.New("could not copy code over to the executable memory area")

// RegisterType defines a list of types we accept and they each fit in a register on x86-64
type RegisterType interface {
	~bool | ~int | ~int32 | ~int64 | ~uint | ~uint32 | ~uint64 | ~float32 | ~float64
}

type buffer struct {
	ptr  unsafe.Pointer
	size int
}

type Jit[I RegisterType, O RegisterType] struct {
	buffers []buffer
}

func (j *Jit[I, O]) Destroy() {
	for _, buf := range j.buffers {
		b := unsafe.Slice((*byte)(buf.ptr), buf.size)
		if err := syscall.Munmap(b); err != nil {
			panic(err)
		}
	}
}

func (j *Jit[I, O]) NewFunc(code []byte) (func(I) O, error) {
	if len(code) == 0 {
		return nil, nil
	}

	i, o := new(I), new(O)
	if unsafe.Sizeof(i) > 8 {
		return nil, fmt.Errorf("input type %T too large to pass in a register", i)
	}
	if unsafe.Sizeof(o) > 8 {
		return nil, fmt.Errorf("output type %T too large to pass in a register", o)
	}

	pageSize := syscall.Getpagesize()
	size := len(code)
	if size%pageSize != 0 {
		size = ((size / pageSize) + 1) * pageSize
	}

	executableArea, err := syscall.Mmap(0, 0, size, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANONYMOUS|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, err
	}

	bytesWritten := copy(executableArea, code)
	if bytesWritten != len(code) {
		return nil, ErrMemCopy
	}

	if err := syscall.Mprotect(executableArea, syscall.PROT_READ|syscall.PROT_EXEC); err != nil {
		return nil, err
	}

	ptr := unsafe.Pointer(&executableArea[0])
	j.buffers = append(j.buffers, buffer{ptr: ptr, size: size})

	return func(i I) O {
		// reinterpret argument as uint64, this results in generic rdi
		// placement for I, thus the JITed code has to perform movq xmm0, rdi
		// if float32 or float64 is necessary
		var raw uint64
		switch v := any(i).(type) {
		case int:
			raw = uint64(v)
		case int32:
			raw = uint64(v)
		case int64:
			raw = uint64(v)
		case uint:
			raw = uint64(v)
		case uint32:
			raw = uint64(v)
		case uint64:
			raw = v
		case bool:
			if v {
				raw = 1
			} else {
				raw = 0
			}
		case float32:
			raw = uint64(math.Float32bits(v))
		case float64:
			raw = math.Float64bits(v)
		default:
			panic("unsupported type")
		}

		ret := C.trampoline(ptr, C.uint64_t(raw))
		t := *new(O)

		// we cheat and do all casting in an unsafe way, since i dont want the
		// performance penalty of reflection + we have a known set of valid
		// types, we always know ret is uint64_t AND T is a runtime type switch
		// on the generic param
		switch any(t).(type) {
		case bool:
			b := ret != 0
			return *(*O)(unsafe.Pointer(&b))
		case int32:
			v := int32(ret)
			return *(*O)(unsafe.Pointer(&v))
		case int64:
			v := int64(ret)
			return *(*O)(unsafe.Pointer(&v))
		case float32:
			v := math.Float32frombits(uint32(ret))
			return *(*O)(unsafe.Pointer(&v))
		case float64:
			v := math.Float64frombits(uint64(ret))
			return *(*O)(unsafe.Pointer(&v))
		default:
			panic("unsupported type")
		}
	}, nil
}
