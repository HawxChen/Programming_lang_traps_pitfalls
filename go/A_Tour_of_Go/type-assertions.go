package main

import "fmt"

func main() {
    var i interface{} = "hello"

    //i.(T): type assertion
    s := i.(string)
    fmt.Println(s)

    s, ok := i.(string)
    fmt.Println(s, ok)

    //When adding the boolean var for checking if interface supports the type,
    //this statement won't trigger panic. It gaves the chance to test
    f, ok := i.(float64)
    fmt.Println(f, ok)

    ff := i.(float64) //panic
    fmt.Println(ff)

    f = i.(float64) // panic
    fmt.Println(f)
}

