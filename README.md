# jit

Go module for executing machine code instructions, without placing them in a file.

Inspired by http://blog.reverberate.org/2012/12/hello-jit-world-joy-of-simple-jits.html

Example usage:

    code := []byte{0xb8, 0x2a, 0x00, 0x00, 0x00, 0xc3}
    jit.Execute(code)

Includes a small machine code interpreter in `cmd/jitrun`. Installation:

    go get -u github.com/xyproto/jit/cmd/jitrun

Example use of `jitrun`:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -
    Stripped source code:
    b8 2a 00 00 00 c3
    
    Source bytes:
    [184 42 0 0 0 195]
    
    The program returned: 42

Or a silent run with `-s`:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -s -
    $ echo $?
    42

---

* Uses `cgo`.
* MIT licensed.
