package main

import (
	"fmt"
	"math"
)

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getDivisors(n int) []int {
	ret := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			ret = append(ret, i)
		}
	}
	return ret
}

func getDivisorsCnt(n int64) int {
	counter := 0
	limit := int(math.Sqrt(float64(n)))
	for i := 1; i <= limit; i++ {
		if n%int64(i) == 0 {
			counter += 2
		}
	}
	return counter
}

func main() {
	//set up the data
	/*
		fmt.Println(getDivisors(1))
		fmt.Println(getDivisors(3))
		fmt.Println(getDivisors(6))
		fmt.Println(getDivisors(10))
		fmt.Println(getDivisors(15))
		fmt.Println(getDivisors(21))
		fmt.Println(getDivisors(28))
	*/

	result := int64(1)
	counter := 1
	for i := 2; counter <= 500; i++ {
		result += int64(i)
		counter = getDivisorsCnt(result)
	}

	fmt.Print("problem 12: for 500 divisiable ", result)
}
