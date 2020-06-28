package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// // from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    doWalk(t, ch)
    close(ch)
}

func doWalk(t *tree.Tree, c chan int) {
    if t == nil {
        return
    }
    doWalk(t.Left, c)
    c <- t.Value
    doWalk(t.Right, c)
}
/*
1 true 1 true
2 true 2 true
3 true 3 true
4 true 4 true
5 true 5 true
6 true 6 true
7 true 7 true
8 true 8 true
9 true 9 true
10 true 10 true
0 false 0 false
*/

func doWalk_Error(t *tree.Tree, c chan int) {
    if t == nil {
        return
    }
    c <- t.Value
    doWalk(t.Left, c)
    doWalk(t.Right, c)
}
/*
10 true 7 true
false
*/

// Same determines whether the trees
// // t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    c1, c2 := make(chan int), make(chan int)
    go Walk(t1, c1)
    go Walk(t2, c2)
    for {
        v1, err1 := <-c1
        v2, err2 := <-c2
        fmt.Println(v1, err1, v2, err2)
        if err1 != err2 || v1 != v2 {
            return false
        }
        if !err1 {
            break
        }
    }
    return true
}

func main() {
    fmt.Println(Same(tree.New(1), tree.New(1)))
}

