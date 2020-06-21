package main

import "fmt"

func main() {
    m := make(map[string]int)

    m["Answer"] = 42
    fmt.Println("The value:", m["Answer"])

    m["Answer"] = 48
    fmt.Println("The value:", m["Answer"])

    delete(m, "Answer")
    fmt.Println("The value:", m["Answer"])

    v, ok := m["Answer"]
    fmt.Println("The value:", v, "Present?", ok)

    fmt.Println(m["ABC"]) //Even the key is not present, the value still returns with 0 because of the definition: map[string]int

    n := make(map[int]string) //Even the key is not present, the value still returns with "" because of the definition: map[int]string
    fmt.Println("Ans:", n[1])
}

