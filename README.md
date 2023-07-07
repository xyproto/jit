# jit [![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/xyproto/jit/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/jit)](https://goreportcard.com/report/github.com/xyproto/jit)

`jit` is a Go module for executing machine code instructions directly, without having to store instructions in an executable file.

It was inspired by this post: [Hello, JIT World: The Joy of Simple JITs](http://blog.reverberate.org/2012/12/hello-jit-world-joy-of-simple-jits.html).

`jit` is also an extremely simple programming language.

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
b8 2a   // mov 2a into the ax register. b8 is the "mov ax" part. 2a is the value.
00 00
00 c3   // return the value in ax
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

## Dependencies

* `cgo` for the `jit` package.
* The [`hexstring`](https://github.com/xyproto/hexstring) module for the `jitrun` command.

## General info

* License: BSD-3
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
