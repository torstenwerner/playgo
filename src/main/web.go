package main

import (
    "fmt"
    "log"
    "net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, 世界")
}

func main() {
    var h Hello
    err := http.ListenAndServe("0.0.0.0:4000", h)
    if err != nil {
        log.Fatal(err)
    }
}
