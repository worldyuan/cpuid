How to Use

package main

import (
        "fmt"
        cpuid "github.com/jeek120/cpuid"
)

func main() {
        ids := [4]uint32{}
        cpuid.Cpuid(&ids, 0)
        fmt.Printf("%d%d%d%d", ids[0], ids[1], ids[2], ids[3])
}
