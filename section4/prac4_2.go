package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var flagvar int

func init() {
	flag.IntVar(&flagvar, "bit", 256, "256 or 384 or 512")
}

func main() {
	var in string

	flag.Parse()

	fmt.Scanf("%s", in)

	switch flagvar {
	case 256:
		data256 := sha256.Sum256([]byte(in))
    	fmt.Printf("%x\n", data256)
	case 384:
		data384 := sha512.Sum384([]byte(in))
        fmt.Printf("%x\n", data384)
	case 512:
		data512 := sha512.Sum512([]byte(in))
        fmt.Printf("%x\n", data512)
	}

}
