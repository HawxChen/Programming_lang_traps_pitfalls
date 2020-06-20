package main

import "fmt"

//Source: https://tour.golang.org/moretypes/10

/*
* The syntax is different from Python!
* We disallow -1 - surrounding index.
*/
func main() {
    s := []int{2, 3, 5, 7, 11, 13}

    s = s[1:4]
    fmt.Println(s)

    s = s[:2]
    fmt.Println(s)

    /*
    ./prog.go:14:7: invalid slice index -1 (index must be non-negative)

    s = s[-1:]
    fmt.Println(s)
    */
}
