package main

import (
	"fmt"
	"math"
)

func Sqrt2(x float64) float64 {
    var z float64 = 1
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*x)
        fmt.Println("%vth: %v", i, z)
    }
    return z
}

func Sqrt_buggy(x float64) float64{
    var z float64 = 1
    for {
        // Everytime, z is the value: 1 quoted from L.18 because := is the initialization statement
        z, t := z - (z*z - x) / (2*x), z 
        fmt.Printf("T:%T V:%v, T:%T V:%v\n", z, z, t, t)
        if math.Abs(t-z) < 1e-6 {
            break
        }
        /*
            T:float64 V:1.475, T:float64 V:1
            T:float64 V:1.475, T:float64 V:1
            T:float64 V:1.475, T:float64 V:1
            T:float64 V:1.475, T:float64 V:1
            T:float64 V:1.475, T:float64 V:1
            T:float64 V:1.475, T:float64 V:1
        */
    }

    return z
}

func Sqrt(x float64) float64{
    z := float64(1)
    var t float64
    for {
        z, t = z - (z*z - x) / (2*x), z
        fmt.Printf("T:%T V:%v, T:%T V:%v\n", z, z, t, t)
        if math.Abs(t-z) < 1e-6 {
            break
        }
    }

    return z
}

func main() {
    fmt.Println(Sqrt(20))
}

