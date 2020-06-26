package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (MyReader) Read(b []byte) (n int, err error) {
    // fill 'A' in b as many as the capacity instead of only lenath
    b = b[:cap(b)]
    // b = b[:len(b)]
    // Error:
    // got byte 0 at offset 1024, want 'A'
    for i := range b {
        b[i] = 'A'
    }
    return cap(b), nil
}

func main() {
    reader.Validate(MyReader{})
}
