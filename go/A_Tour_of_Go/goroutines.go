package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(250 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world") // Init another threads
    say("hello")
}

