package main

import "fmt"

/*
* map's value of type is constant and only one.
* */

type Vertex struct {
    Lat, Long float64
}

type Vertex2 struct {
    Lat, Long float64
}
//Error: ./prog.go:17:2: cannot use Vertex2 literal (type Vertex2) as type Vertex in map value
var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex2{
        37.42202, -122.08408,
    },
}

func main() {
    fmt.Println(m)
}
