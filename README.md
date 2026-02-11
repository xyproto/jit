# jit [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/jit/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/jit)](https://goreportcard.com/report/github.com/xyproto/jit)

`jit` is a Go module for executing machine code instructions directly, without having to store instructions in an executable file.

It was inspired by this post: [Hello, JIT World: The Joy of Simple JITs](http://blog.reverberate.org/2012/12/hello-jit-world-joy-of-simple-jits.html).

`jit` is also an extremely simple programming language.

For a similar project written in Rust, check out [machinecode](https://github.com/xyproto/machinecode).

## Example usage

```go
code := []byte{0xb8, 0x2a, 0x00, 0x00, 0x00, 0xc3}
jit.Execute(code)
```

## The `jitrun` command

`jit` includes a small machine code execution program named `jitrun` that can be installed with Go 1.17 or later:

```bash
go install github.com/xyproto/jit/cmd/jitrun@latest
```

You can provide a file, and run it:

*`42.mc`*:

```
// This program moves 42 into eax and returns

b8 2a 00 00 00  // mov 2a000000 into the eax register. b8 is the "mov eax" part. 0x2a is 42.
c3              // return to the caller (the return value is held in eax)
```

    ./jitrun 42.mc

It's possible to pipe the machine code directly to `jitrun`:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -
    Stripped source code:
    b8 2a 00 00 00 c3

    Source bytes:
    [184 42 0 0 0 195]

    The program returned: 42

`jitrun` can be run silently with the `-s` flag:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -s -
    $ echo $?
    42

Here is another example program:

```
// This program takes the square root of 1024 and returns the answer (in eax), which is 32

b8 00 04 00 00  // mov 1024 (0x400) into eax
f3 0f 2a c0     // mov eax into the xmm0 register
f3 0f 51 c0     // take the square root of the xmm0 register and place it into xmm0
f3 0f 2c c0     // move xmm0 back into eax
c3              // return to the caller (the return value is held in eax)
```

## Dependencies

* `cgo` for the `jit` package.
* The [`hexstring`](https://github.com/xyproto/hexstring) module for the `jitrun` command.

## General info

* License: BSD-3
* Version: 1.0.1
