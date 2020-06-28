package main

import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // send sum to c
}

func main() {
    s := []int{7, 2, 8, -9, 4, 0}

    c := make(chan int)
    go sum(s[:len(s)/2], c) //random ordr
    go sum(s[len(s)/2:], c) //random order
    x, y := <-c, <-c // receive from c

    fmt.Println(x, y, x+y)
}

// Result could be
// -5 17 12
// 17 -5 12
