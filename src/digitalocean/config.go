package main

import "os"

type Config struct {
    Token string
}

func Fetch() (*Config) {
    token := os.Getenv("DO_TOKEN")
    if (token == "") {
        panic("environment variable DO_TOKEN not set")
    }
    return &Config{token}
}