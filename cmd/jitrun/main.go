package main

// This is a machine code interpreter

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/xyproto/hexstring"
	"github.com/xyproto/jit"
)

func usageExit() {
	fmt.Fprintln(os.Stderr, "Please provide the filename of a hex encoded machine code program as the first argument.")
	fmt.Fprintln(os.Stderr, "Use -s to silence the output and only set the exit code to the return code of the program.")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usageExit()
	}
	filename := os.Args[1]
	verbose := true

	// Silent flag?
	if os.Args[1] == "-s" || os.Args[1] == "--silent" {
		if len(os.Args) < 3 {
			usageExit()
		}
		filename = os.Args[2]
		verbose = false
	}

	// Read the source file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	// Strip comments (starting with ";", "#" or "//")
	var sourceCode strings.Builder
	for i, line := range strings.Split(string(data), "\n") {
		if i > 0 {
			sourceCode.WriteString("\n")
		}
		spos := strings.Index(line, ";")
		if spos != -1 {
			line = line[:spos]
		}
		hpos := strings.Index(line, "#")
		if hpos != -1 {
			line = line[:hpos]
		}
		cpos := strings.Index(line, "//")
		if cpos != -1 {
			line = line[:cpos]
		}
		sourceCode.WriteString(strings.TrimSpace(line))
	}

	if verbose {
		fmt.Printf("Stripped source code:\n%s\n", sourceCode.String())
	}

	code, err := hexstring.StringToBytes(sourceCode.String())
	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Printf("Source bytes:\n%v\n\n", code)
	}

	retval, err := jit.Execute(code)
	if err != nil {
		panic(err)
	}

	if verbose {
		fmt.Printf("The program returned: %d\n", retval)
	} else {
		os.Exit(retval)
	}
}
