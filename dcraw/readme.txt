This package contains the following files:

1) dcraw.exe (32-bit version)
2) dcraw64.exe (64-bit version)
3) dcraw.c (source code by Dave Coffin, https://www.dechifro.org/dcraw/)
4) readme.txt (this file)

dcraw.c was compiled with full optimization (using the Microsoft C/C++ optimizing compiler). The two executables are digitally signed.

Full history available at https://github.com/ncruces/dcraw

--

Linux copy compiled with "gcc -o dcraw -O4 dcraw.c -lm -DNODEPS", 64 bit.

32 bit windows compiled with "i686-w64-mingw32-gcc -o dcraw.exe -O4 dcraw.c -lm -DNODEPS -L/usr/i686-w64-mingw32/lib/ -lwsock32"

64 bit windows compiled with "x86_64-w64-mingw32-gcc -o dcraw64.exe -O4 dcraw.c -lm -DNODEPS -L/usr/x86_64-w64-mingw32/lib/ -lwsock32"

