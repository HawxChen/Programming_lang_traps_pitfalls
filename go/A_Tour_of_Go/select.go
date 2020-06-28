package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select { // Don't confuse against switch case block
        case c <- x: // waiting goroutine
            x, y = y, x+y
        case <-quit: // waiting goroutine
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- nil // The practical pattern ends the infinit loop in main thread
    }()
    fibonacci(c, quit)
}
