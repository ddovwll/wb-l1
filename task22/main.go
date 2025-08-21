package main

import (
	"fmt"
	"math"
	"math/big"
)

func Add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func Mul(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func main() {
	a := big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64))
	b := big.NewInt(0).Mul(big.NewInt(math.MaxInt64), big.NewInt(math.MaxInt64/2))
	fmt.Println(a)
	fmt.Println(b)

	println()

	fmt.Println(Add(a, b))
	fmt.Println(Sub(a, b))
	fmt.Println(Mul(a, b))
	fmt.Println(Div(a, b))
}
