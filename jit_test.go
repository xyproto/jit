package jit

import (
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

	expected := int32(64 * 64)
	got := function(64)
	if got != expected {
		t.Fatalf("got(%v) != expected(%v)", got, expected)
	}
	j.Destroy()
}
