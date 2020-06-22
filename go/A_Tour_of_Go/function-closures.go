package main

import "fmt"

func adder() func(int) int {
    sum := 0 // How it becomes static varable!?
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    pos, neg := adder(), adder()
    for i := 0; i < 10; i++ {
        fmt.Println(
            pos(i),
            neg(-2*i),
        )
    }
    fmt.Println(pos(1), neg(-1))
}
/*
* 0 0
* 1 -2
* 3 -6
* 6 -12
* 10 -20
* 15 -30
* 21 -42
* 28 -56
* 36 -72
* 45 -90
* 46 -91
*/
