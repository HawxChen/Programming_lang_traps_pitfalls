// Source: https://tour.golang.org/methods/6

/*
var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
*/

/*
For the statement v.Scale(5), even though v is a value and not a pointer,
the method with the pointer receiver is called automatically.
That is, as a convenience.
Go interprets the statement v.Scale(5) as (&v).Scale(5) 
ince the Scale method has a pointer receiver.
*/
/*
var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
*/

package main

import "fmt"

type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func (v Vertex) Abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Abs() float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
      return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
    v := Vertex{3, 4}

    // automatically interpret v.Scale(5) as (&v).Scale(5) in Go.
    v.Scale(2)

    ScaleFunc(&v, 10)

    p := &Vertex{4, 3}

    // pass by value, won't change p
    // In this case,
    // the method call p.Abs() is Automatically interpreted as (*p).Abs().
    p.Abs()

    p.Scale(3)
    ScaleFunc(p, 8)
    AbsFunc(*p, 8)

    fmt.Println(v, p)
}
