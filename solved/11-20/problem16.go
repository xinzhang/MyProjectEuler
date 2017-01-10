package main

import "fmt"
import "math/big"
import "strconv"

func Power(n int64, pow int) *big.Int {
	r := big.NewInt(1)
	result := big.NewInt(1)
	for i := 0; i < pow; i++ {
		result = r.Mul(big.NewInt(n), result)
	}
	return result
}

func main() {
	val := Power(2, 1000)
	fmt.Println(val)
	result := 0
	r := []rune(val.String())
	for _, v := range r {
		i, _ := strconv.Atoi(string(v))
		result += i
	}
	fmt.Println(result)
}
