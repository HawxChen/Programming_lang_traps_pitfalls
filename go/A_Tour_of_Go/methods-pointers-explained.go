// Source: https://tour.golang.org/methods/5


package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// pass by value
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// pass by  address
func Scale(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

func main() {
    v := Vertex{3, 4}
    Scale(&v, 10)
    fmt.Println(Abs(v))
}
