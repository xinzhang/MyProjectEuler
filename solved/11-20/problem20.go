package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func calcTotalDigit(x *big.Int) int {
	result := 0
	r := []rune(x.String())
	for _, v := range r {
		i, _ := strconv.Atoi(string(v))
		result += i
	}
	return result
}

func main() {
	v := factorial(big.NewInt(100))
	result := calcTotalDigit(v)
	fmt.Println(result)
}
