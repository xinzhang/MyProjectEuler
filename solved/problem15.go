package main

import (
	"fmt"
	"math/big"
)

// func factorial(n uint64) uint64 {
// 	if n == 1 {
// 		return uint64(1)
// 	} else {
// 		return uint64(n) * factorial(n-1)
// 	}
// }
//
// // (Private) Returns the tuple (F(n), F(n+1)).
// func fib(n int64) *big.Int {
// 	if n == 1 {
// 		return big.NewInt(1)
// 	} else {
// 		return big.NewInt(0).Mul(big.NewInt(n), fib(n))
// 	}
// }

//lattice paths

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}

func possibility(n1 int64, n2 int64) *big.Int {
	n := big.NewInt(1)
	bn1 := big.NewInt(n1)
	bn2 := big.NewInt(n2)
	bn12 := big.NewInt(n1 - n2)

	k := n.Mul(factorial(bn2), factorial(bn12))
	fmt.Println(k)
	fmt.Println(factorial(bn1))
	return n.Div(factorial(bn1), k)
}

func main() {
	//fmt.Println(fib(40))
	//r := big.NewInt(40)
	//fmt.Println(factorial(r))
	//fmt.Println(fib(20) * fib(20))

	fmt.Println(possibility(4, 2))
	fmt.Println(possibility(6, 3))
	fmt.Println(possibility(40, 20))
}
