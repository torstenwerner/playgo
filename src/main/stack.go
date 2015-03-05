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

func printFact(n uint64) {
    fact := factorial(n)
    bitLen := fact.BitLen()
    ratio := float32(bitLen) / float32(n)
    fmt.Printf("%5d %6d %4.1f %d\n", n, bitLen, ratio, fact)
}

func main() {
    for n := uint64(1); n <= 40000; n = n * 2 {
        printFact(n)
    }
}