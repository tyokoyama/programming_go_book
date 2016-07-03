// 練習問題1.3
package section1

import (
    "strings"
)

func Join(args []string) string {
    return strings.Join(args, " ")
}

func JoinBasic(args []string) string {
    var s, sep string
    for i := 0; i < len(args); i++ {
        s += sep + args[i]
        sep = " "
    }

    return s
}