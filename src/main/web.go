package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
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

type Greeting struct {
    Number int
    Name string
}

type Hello struct {
    ch <-chan int
}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Path[1:]
    if name == "" {
        name = "World"
    }
    greeting := Greeting{<-h.ch, name}
    asJson, _ := json.Marshal(greeting)
    fmt.Fprintf(w, "%s", asJson)
}

func count(ch chan<- int) {
    counter := 0
    for {
        ch <- counter
        counter ++
    }
}
