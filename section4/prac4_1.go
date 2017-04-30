package main
import (
    "crypto/sha256"
    "fmt"
)

var pc1 [256]byte
var pc2 [256]byte

func main()  {
    for i := range pc1 {
        pc1[i] = pc1[i/2] + byte(i&1)
    }

    for i := range pc2 {
        pc2[i] = pc2[i/2] + byte(i&1)
    }

    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

    fmt.Printf("Count = %d\n", PopCount(c1, c2))
}

func PopCount(c1, c2 [32]byte) int {

    bitCount := 0

    // // とりあえず力づくで計算する方法。
    // for i := 0; i < len(c1); i++ {
    //     for j := 0; j < 8; j++ {
    //         if ((c1[i] >> uint8(j)) & 0x01) != (uint8(c2[i] >> uint8(j)) & 0x01) {
    //             bitCount++
    //         }
    //     }
    // }

    // 論理演算（EX-OR）^を使う方法
    for i := 0; i < len(c1); i++ {
        bt := (c1[i] ^ c2[i])
        for j := 0; j < 8; j++ {
            if ((bt >> uint8(j)) & 0x01) == 1 {
                bitCount++
            }
        }
    }

    return bitCount
}