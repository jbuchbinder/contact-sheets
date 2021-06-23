all: clean linux windows-32 windows-64

clean:
	rm -f dcraw dcraw*exe

linux:
	gcc -o dcraw -O4 dcraw.c -lm -DNODEPS

windows-32:
	i686-w64-mingw32-gcc -o dcraw.exe -O4 dcraw.c -lm -DNODEPS -L/usr/i686-w64-mingw32/lib/ -lwsock32

windows-64:
	x86_64-w64-mingw32-gcc -o dcraw64.exe -O4 dcraw.c -lm -DNODEPS -L/usr/x86_64-w64-mingw32/lib/ -lwsock32

