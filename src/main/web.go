package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    ch := make(chan int, 3)
    go count(ch)
    h := Hello{ch}
    fmt.Println("application started")
    err := http.ListenAndServe("0.0.0.0:4000", h)
    if err != nil {
        log.Fatal(err)
    }
}

type Hello struct {
    ch <-chan int
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "%d: Hello, 世界", h.getCounter())
}

func (h Hello) getCounter() int {
    return <-h.ch
}

func count(ch chan <- int) {
    counter := 0
    for {
        ch <- counter
        counter ++
    }
}
