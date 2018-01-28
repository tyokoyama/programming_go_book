package main

/*
#include <stdio.h>
void wrapprintf(char* s) {
	printf("%s", s);
}
*/
import "C"

func main() {
	s := C.CString("Hello World\n")
	C.wrapprintf(s)
}
