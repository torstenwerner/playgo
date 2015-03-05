package main

import (
    "fmt"
    "math/big"
)

func factorial(n uint64) *big.Int {
    bigN := new(big.Int).SetUint64(n)
    switch n {
        case 0, 1, 2: return bigN;
    }
    return bigN.Mul(bigN, factorial(n - 1))
}

func main() {
    fact := factorial(10000)
    fmt.Println(fact.BitLen())
    fmt.Println(fact)
}