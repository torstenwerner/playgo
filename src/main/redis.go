package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    conn, err := redis.Dial("tcp", "192.168.59.103:6379")
    if err != nil {
        panic(err)
    }
    print(conn, "Set", "hello", 26)
    print(conn, "Incr", "hello")
    print(conn, "Get", "hello")
}

func print(conn redis.Conn, commandName string, args ...interface{}) {
    value, err := conn.Do(commandName, args...)
    if (err != nil) {
        panic(err)
    }
    switch v := value.(type) {
        case string: fmt.Printf("got %s\n", v)
        case int64: fmt.Printf("got %d\n", v)
        case []uint8: fmt.Printf("got %s\n", string(v))
        default: fmt.Printf("got unknown %s type %T\n", v, v)
    }
}
