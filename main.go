package main

import (
        "fmt"
        //"os"
        "strconv"
)

func TheArbitraryBaseConverter(bin uint8, bout uint8, s string) string {
        n, _ := strconv.ParseUint(s, int(bin), 64)
        return strconv.FormatUint(n, int(bout))
}

func main() {
        fmt.Println(TheArbitraryBaseConverter(27, 10, "1m"))
}
