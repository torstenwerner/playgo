package main

import "fmt"

func main() {
    h := House{n: "tent"}
    fmt.Println(h.name())
}

type House struct {
    n string
}

func (h House) name() (n string) {
    if (h.n == "") {
        h.n= "doctor"
    }
    n = h.n
    return
}