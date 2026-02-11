# hexstring

Deal with strings of space separated hexadecimal bytes.

## Example use

Convert a string with space separated numbers from `0x00` to `0xff` to a byte slice:

```go
helloWorldBytes, err := StringToBytes("48 65 6c 6c 6f 20 57 6f 72 6c 64 21")
```

## General info

* License: MIT
* Version: 1.0.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
