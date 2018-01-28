package main

/*
#cgo LDFLAGS: -lm
#include <math.h>
*/
import "C"

import (
	"fmt"
)

func main() {
	s, _ := C.sqrt(2)

	fmt.Println(s)
}
