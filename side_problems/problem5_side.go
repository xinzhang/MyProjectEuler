package main

import "fmt"

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IsPrime(x int) bool {
	for y := 2; y <= x/2; y++ {
		if x%y == 0 {
			return false
		}
	}
	return true
}

func AllMutliply(n int) int {
	result := 1
	for i := 2; i < n+1; i++ {
		if IsPrime(i) {
			fmt.Print(i)
			result = result * i
		}
	}
	return result
}

func main() {
	/* test prime number
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))
	*/

	fmt.Println(AllMutliply(10))

	//y := AllMutliply(20)
	//fmt.Println("the smallest positive number that is evenly divisible by all of the numbers from 1 to 20 ", y)
}
