// This program takes the square root of 1024 and returns the answer (in eax), which is 32

b8 00 04 00 00  // mov 1024 (0x0400) into eax (0x00 comes first and then 0x04, because it is little-endian)
f3 0f 2a c0     // mov eax into the xmm0 register
f3 0f 51 c0     // take the square root of the xmm0 register and place it into xmm0
f3 0f 2c c0     // move xmm0 back into eax
c3              // return to the caller (the return value is held in eax)
