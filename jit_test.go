package jit

import (
	"testing"

	"github.com/xyproto/hexstring"
)

func TestExecution(t *testing.T) {
	machineCodeProgram := `
b8 00
00 00
00 c3
`

	code, err := hexstring.StringToBytes(machineCodeProgram)
	if err != nil {
		t.Fatal(err)
	}

	// Change eax in the above program to be set to a value
	num := 21
	code[1] = uint8(num)

	retval, err := Execute(code)
	if err != nil {
		t.Fatal(err)
	}

	if retval != num {
		t.Fatalf("The wrong value was returned: %d != %d\n", retval, num)
	}
}
