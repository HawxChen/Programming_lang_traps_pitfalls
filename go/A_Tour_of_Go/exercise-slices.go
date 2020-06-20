package main
import (
    "golang.org/x/tour/pic"
)
func Pic(dx, dy int) [][]uint8 {
    ss := make([][]uint8, dy)
    for y := range ss {
        s := make([]uint8, dx)
        for x := range s{
            s[x] = uint8((x+y)/2)
        }
        ss[y] = s //unlike C.
        /*
        * We need to know "=", assignment operator's difference between Python and C
        * uint8[] = uint8[]
        * C - point assignment to only an element
        * Python - object assignment to only an element
        * Go - memory copy in this case.
        */
    }
    return ss
}


func main() {
    pic.Show(Pic)
}
