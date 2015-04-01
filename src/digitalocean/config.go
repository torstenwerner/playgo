package main

import (
    "io/ioutil"
    "os"
    "gopkg.in/yaml.v2"
)

type Config struct {
    Token string
}

func Fetch() (*Config) {
    home := os.Getenv("HOME")
    if (home == "") {
        panic("environment variable HOME not set")
    }

    filename := home + "/.digitalocean.yml"
    input, err := ioutil.ReadFile(filename)
    if (err != nil) {
        panic(err)
    }

    var config Config
    yaml.Unmarshal(input, &config)

    return &config
}