package main

//#include <stdlib.h>
//#include "./lib/lib.h"
import "C"
import "unsafe"

func main() {
	cs := C.CString("hello, world!")
	defer C.free(unsafe.Pointer(cs))
	C.myPrint(cs)
}
