package main

import (
    "golang.org/x/tour/wc"
    "strings"
    "fmt"
)

func WordCount(s string) map[string]int {
    tokens := strings.Fields(s)
    m := make(map[string]int)

    //to print list needs format identifier.
    fmt.Printf("tokens: %q\n", tokens)
    for idx := range tokens {
        m[tokens[idx]] += 1
    }

    return m
}

func main() {
    wc.Test(WordCount)
}

