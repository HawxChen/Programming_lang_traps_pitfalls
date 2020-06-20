// Source https://tour.golang.org/moretypes/14
package main

import "fmt"

/*
*   1. Use C thinking.
*
*/


func main() {
    a := make([]int, 5)
    printSlice("a", a)

    b := make([]int, 0, 5)
    printSlice("b", b)

    c := b[:2]
    printSlice("c", c)

    d := c[2:5]
    printSlice("d", d)

    printSlice("a", a)  // Doesn't harm original a
    printSlice("b", b)  // Doesn't harm original b
}
/*
    a len=5 cap=5 [0 0 0 0 0]
    b len=0 cap=5 []
    c len=2 cap=5 [0 0]
    d len=3 cap=3 [0 0 0]
    a len=5 cap=5 [0 0 0 0 0]
    b len=0 cap=5 []
*/

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n",
    s, len(x), cap(x), x)
}
