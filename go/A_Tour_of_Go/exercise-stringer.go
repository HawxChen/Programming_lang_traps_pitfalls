ckage main

import (
    "fmt"
    "strings"
    "strconv"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr

// String() string - stringers is similar to __repr__ in Python
func (ip IPAddr) String() string {
    //make is for the memory of container allocation
    str_array := make([]string, len(ip))
    for i, v := range ip {
        str_array[i] = strconv.Itoa(int(v))
    }
    return strings.Join(str_array, ".")
}

func main() {
    hosts := map[string]IPAddr{
        "loopback":  {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
