package main

import (
    "fmt"
    "os"
    "strconv"
)

func main()  {
    for pos, arg := range os.Args {
        s := strconv.Itoa(pos) + " " + arg
        fmt.Println(s)        
    }
}