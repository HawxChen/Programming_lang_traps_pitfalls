// Source: https://tour.golang.org/methods/11


package main

import (
    "fmt"
    "math"
)

type Abser interface {
    Abs() float64
}

type MyFloat float64

/*
* This method means type MyFloat implements the interface Abser,
* but we don't need to explicitly declare that it does so.
*/
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

type Vertex struct {
    X, Y float64
}

/*
* This method means type *Vertex implements the interface Abser,
* but we don't need to explicitly declare that it does so.
*/
/*
* Calling a method on an interface value executes the method of the same name on its underlying type.
*/
func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    var a Abser
    f := MyFloat(-math.Sqrt2)
    v := Vertex{3, 4}

    a = f  // a MyFloat implements Abser
    a = &v // a *Vertex implements Abser
    v.Abs() // worked! automatically interpret v.Abs() as (&v).Abs()
    (&v).Abs() // worked.
    a.Abs() // worked because the type of a is *Vertex.

    // In the following line, v is a Vertex (not *Vertex)
    // and does NOT implement Abser.
    // TBD: more detailed explanation.
    a = v
    //Error: 
    //./prog.go:43:4: cannot us
    //v (type Vertex) as type Abser in assignment:
    //  Vertex does not implement Abser (Abs method has pointer receiver)

    fmt.Println(a.Abs())
}
