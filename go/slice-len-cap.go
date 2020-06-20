//Source: https://tour.golang.org/moretypes/11

/*
* 1. Use C thinking. Focus on s' inherited status changes.
* 2. := is an initialization operator in Go,
*    which does not exist in Python
* 3. = is the reference, pointer assignment,
*    not like python doing whole memory/object replacement.
*      Original instances in Python have been replaced via 'assignment operator'
     * Python
        >>> x = [1, 2, 3, 4, 5]
        >>> >>> x
        >>> >>> [1, 2, 3, 4, 5]
        >>> >>> >>> x = x[1:4]
        >>> >>> >>> >>> x
        >>> >>> >>> >>> [2, 3, 4]
* 4. len - length
* 5. cap - capacity does not exist in python
*/

package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    // Slice the slice to give it zero length.
    s = s[:0]
    printSlice(s)

    // Extend its length.
    s = s[:4]
    printSlice(s)

    // Drop its first two values.
    s = s[2:]
    printSlice(s)

    /*
    len=6 cap=6 [2 3 5 7 11 13]
    len=0 cap=6 []
    len=4 cap=6 [2 3 5 7]
    len=2 cap=4 [5 7]
    */
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

