whyckage main

import "fmt"

// fibonacci is a function that returns
// // a function that returns an int.
func fibonacci() func() int {
    f1, f2 := -1, -1

    return func() int{ //nested func has no name.
        if f1 == -1 {
            f1 = 1
            return 0
        }
        if f2 == -1 {
            f2 = 0
            return 1
        }
        f1, f2 = f1 + f2, f1
        return f1
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
