// Q: https://tour.golang.org/methods/23
package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (rot *rot13Reader) Read (b []byte) (int, error) {
    len, err := rot.r.Read(b)
    for i := 0; i < len; i++ {
        if (b[i] >= 'A' && b[i] < 'N') || (b[i] >= 'a' && b[i] < 'n') {
            b[i] += 13
        } else if (b[i] > 'M' && b[i] <= 'Z') || (b[i] > 'm' && b[i] <= 'z') {
            b[i] -= 13
        }
    }

    return len, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

