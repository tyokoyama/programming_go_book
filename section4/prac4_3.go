package main 

import (
    "fmt"
)

func main() {
    data := [...]int{0, 1, 2, 3, 4, 5}

    reverse(&data)

    fmt.Println(data)
}

func reverse(data *[6]int) {
    for i,j := 0, len(data) - 1; i < j; i, j = i+1, j-1 {
        data[i], data[j] = data[j], data[i]
    }
}