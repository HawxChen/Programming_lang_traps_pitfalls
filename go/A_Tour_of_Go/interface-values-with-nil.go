// Source: https://tour.golang.org/methods/12

// In Go it is common to write methods that gracefully handle being called with 
// a nil receiver (as with the method M in this example.)

package main

import "fmt"

type I interface {
    M()
}

type T struct {
    S string
}

func (t *T) M() {
    //MUST handle NIL case
    if t == nil {
        fmt.Println("<nil>")
        return
    }
    fmt.Println(t.S)
}

func main() {
    var i I

    // No memory created
    // The pointer to NIL
    var t *T

    //Pointer assignment
    i = t
    describe(i)
    i.M()

    i = &T{"hello"}
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}

