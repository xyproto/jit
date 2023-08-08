// This program moves 42 into eax and returns

b8 2a 00 00 00  // mov 0x2a000000 into the eax register. b8 is the "mov eax" part. 2a is 42.
c3              // return to the caller (the return value is held in eax)
