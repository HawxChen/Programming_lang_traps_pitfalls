package main

import "fmt"
/*
* The syntax is similar to Ptyhon; however, memory management is TOTALLY different. 
*/

//Source: https://tour.golang.org/moretypes/8

/*
Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
*/
func main() {
    names := [4]string{
        "John",
        "Paul",
        "George",
        "Ringo",
    }
    fmt.Println(names)

    a := names[0:2] // share the same memory of names
    b := names[1:3] // share the same memory of names
    fmt.Println(a, b)

    b[0] = "XXX" // share the same memory of names
    fmt.Println(a, b)
    fmt.Println(names)
    /*
    [John Paul George Ringo]
    [John Paul] [Paul George]
    [John XXX] [XXX George]
    */
}

