package main

import (
	"fmt"
	"github.com/xyproto/hexstring"
	"github.com/xyproto/jit"
	"log"
	"math/rand"
	//"os"
	//"os/signal"
	"runtime/debug"
	//"syscall"
	"time"
)

// findCode will generate a slice of machine code instructions with length
// codeLength that returns the targetValue, by random trial and error.
// The code is ran and may do anything. Beware!
func findCode(codeLength, targetValue int) []byte {

	// Set up a signal handler for illegal instructions (SIGILL)
	// and for segfaults (SIGSEGV)
	//sigChan := make(chan os.Signal, 1)
	//signal.Notify(sigChan, syscall.SIGILL, syscall.SIGSEGV)
	//go func() {
	//	for range sigChan {
	//		fmt.Println("Illegal instruction, or segfault.")
	//	}
	//}()

	debug.SetPanicOnFault(true)

	// Set up a panic handler
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered 1")
		}
	}()

	code := make([]byte, codeLength)

OUT:
	for {
		// Create random code
		rand.Read(code)

		fmt.Printf("Running code [%v] ...", hexstring.BytesToString(code))

		retvalChan := make(chan int, 1)

		// Run the machine code in a goroutine
		go func() {

			debug.SetPanicOnFault(true)

			// Set up a panic handler
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered 2")
				}
			}()

			// Run the code
			retval, err := jit.Execute(code)
			if err != nil {
				retvalChan <- -1
			} else {
				retvalChan <- retval
			}
		}()

		// Listen to the channels
		select {
		case retval := <-retvalChan:
			if retval == targetValue {
				log.Println("returned value is correct!")
				break OUT
			} else {
				log.Println("returned value is wrong.")
				continue OUT
			}
		case <-time.After(1 * time.Second):
			fmt.Println("timeout")
			continue OUT
		}

		// Clear the slice
		code = code[:0]
	}

	return code
}

func main() {
	fmt.Printf("Looking for 6 bytes of machine code that returns 42.\n\n")
	code := findCode(6, 42)
	fmt.Println("Success! Found machine code that returns 42:")
	fmt.Println(hexstring.BytesToString(code))
}
