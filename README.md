# jit [![Build Status](https://travis-ci.org/xyproto/jit.svg?branch=master)](https://travis-ci.org/xyproto/jit) [![GoDoc](https://godoc.org/github.com/xyproto/jit?status.svg)](https://godoc.org/github.com/xyproto/jit) [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/jit/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/jit)](https://goreportcard.com/report/github.com/xyproto/jit)

`jit` is a Go module for executing machine code instructions directly, without having to store instructions in an executable file.

It was inspired by this post: [Hello, JIT World: The Joy of Simple JITs](http://blog.reverberate.org/2012/12/hello-jit-world-joy-of-simple-jits.html).

## Example usage

```go
code := []byte{0xb8, 0x2a, 0x00, 0x00, 0x00, 0xc3}
jit.Execute(code)
```

## The `jitrun` command

`jit` includes a small machine code execution program nameed `jitrun` that can be installed with:

```bash
go get -u github.com/xyproto/jit/cmd/jitrun
```

Example use of `jitrun`:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -
    Stripped source code:
    b8 2a 00 00 00 c3

    Source bytes:
    [184 42 0 0 0 195]

    The program returned: 42

Alternatively, provide a file, and run it:

*`42.mc`*:

```
b8 2a   // mov 2a into the ax register. b8 is the "mov ax" part. 2a is the value.
00 00
00 c3   // return the value in ax
```

    ./jitrun 42.mc

`jitrun` can also be run silently by using the `-s` flag:

    $ echo 'b8 2a 00 00 00 c3' | jitrun -s -
    $ echo $?
    42

## Dependencies

The only dependencies are `cgo` (for the `jit` package) and the [`hexstring`](https://github.com/xyproto/hexstring) module (only for the `jitrun` command).

## General info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
