package main

import "fmt"

func IsPrime(x int) (bool, int) {
	for y := 2; y < x/2; y++ {
		if x%y == 0 {
			return false, y
		}
	}
	return true, 0
}

func GetMaxFactor(x int) int {
	for y := 2; y < x/2; y++ {
		if x%y == 0 {
			return GetMaxFactor(x / y)
		}
	}
	return x
}

func main() {
	fmt.Println(GetMaxFactor(100))
	fmt.Println(GetMaxFactor(111))
	fmt.Println(GetMaxFactor(29))
	fmt.Println(GetMaxFactor(13195))

	var result int = 0
	var theNumber int = 600851475143
	result = GetMaxFactor(theNumber)

	fmt.Println("The largest prime factor of ", theNumber, " is ", result)
}
