package main

import "C"
import "log"

//export myPrint
func myPrint(str *C.char) {
	gStr := C.GoString(str)
	log.Println(gStr)
}

func main() {}
