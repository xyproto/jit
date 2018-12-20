# jit

Go module for executing code from memory.

Inspired by http://blog.reverberate.org/2012/12/hello-jit-world-joy-of-simple-jits.html

Example usage:

    code := []byte{0xb8, 0x2a, 0x00, 0x00, 0x00, 0xc3}
    jit.Execute(code)

Includes a small machine code interpreter in `cmd/jitrun`. Installation:

    go get -u github.com/xyproto/jit/cmd/jitrun

---

* Uses `cgo`.
* MIT licensed.
