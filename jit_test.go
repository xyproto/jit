package jit

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSquare(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("amd64 only")
	}
	code := []byte{
		0x48, 0x89, 0xF8, // mov rax, rdi
		0x48, 0x0F, 0xAF, 0xC0, // imul rax, rax
		0xC3, // ret
	}

	j := Jit[int32, int32]{}
	function, err := j.NewFunc(code)
	if err != nil {
		t.Fatalf("jit failed: %v", err)
	}
	fmt.Printf("%p\n", function)

	expected := int32(64 * 64)
	got := function(64)
	if got != expected {
		t.Fatalf("got(%v) != expected(%v)", got, expected)
	}
	j.Destroy()
}

func TestExecute_ReturnConstant(t *testing.T) {
	if runtime.GOARCH != "amd64" {
		t.Skip("amd64 only")
	}

	tests := []struct {
		name string
		val  int32
	}{
		{"zero", 0},
		{"small", 21},
		{"large", 123456789},
		{"negative", -42},
		{"max", 0x7fffffff},
		{"min", -0x80000000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code := make([]byte, 6)

			// mov eax, imm32
			code[0] = 0xB8

			// ret
			code[5] = 0xC3

			j := Jit[int32, int32]{}
			function, err := j.NewFunc(code)
			if err != nil {
				t.Fatalf("jit failed: %v", err)
			}

			got := function(tt.val)
			j.Destroy()
			if int32(got) != tt.val {
				t.Fatalf("wrong return: got %d, want %d", got, tt.val)
			}
		})
	}
}
