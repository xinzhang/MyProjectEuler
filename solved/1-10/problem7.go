package main

import "fmt"

func IsPrime(x int) bool {
	for y := 2; y <= x/2; y++ {
		if x%y == 0 {
			return false
		}
	}
	return true
}

func GetNthPrime(x int) int {
	cnt, startVal := 0, 1
	for {
		startVal++
		if IsPrime(startVal) {
			cnt++
		}

		if cnt == x {
			return startVal
		}
	}
}

func main() {
	fmt.Println(GetNthPrime(1))
	fmt.Println(GetNthPrime(2))
	fmt.Println(GetNthPrime(3))
	fmt.Println(GetNthPrime(6))
	fmt.Println(GetNthPrime(10001))
}
