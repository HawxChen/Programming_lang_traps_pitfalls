package main

import "fmt"

func main() {
    ch := make(chan int, 1)
    ch <- 1
    // ch <- 2 // if we comment this line, the following error happens
    // fatal error: all goroutines are asleep - deadlock!
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

func main() {
    ch := make(chan int, 2)
    ch <- 1
    ch <- 2 // if we comment this line, the following error happens
    // fatal error: all goroutines are asleep - deadlock!
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

