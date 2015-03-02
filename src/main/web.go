package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    ch := make(chan int, 3)
    go count(ch)
    server := Hello{ch}
    fmt.Println("application started")
    err := http.ListenAndServe("0.0.0.0:4000", server)
    if err != nil {
        log.Fatal(err)
    }
}

type Hello struct {
    ch <-chan int
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[1:]
    if name == "" {
        name = "World"
    }
    fmt.Fprintf(w, "Greeting #%d: Hello, %s!", <-h.ch, name)
}

func count(ch chan<- int) {
    counter := 0
    for {
        ch <- counter
        counter ++
    }
}
