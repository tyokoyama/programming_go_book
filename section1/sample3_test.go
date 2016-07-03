//
package section1

import (
    "fmt"
    "strconv"
    "testing"
)

func BenchmarkJoin(b *testing.B) {
    var args []string
    for i := 0; i < 10000; i++ {
        args = append(args, strconv.Itoa(i))
    }
    fmt.Println("Join")
//    fmt.Println(Join(args))
    Join(args)
}

func BenchmarkJoinBasic(b *testing.B) {
    var args []string
    for i := 0; i < 10000; i++ {
        args = append(args, strconv.Itoa(i))
    }
    fmt.Println("JoinBasic")
//    fmt.Println(JoinBasic(args))
    JoinBasic(args)
}