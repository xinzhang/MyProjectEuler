package main

import "fmt"

func IsFactor(x int) (bool, int) {
	for y := 2; y < x/2; y++ {
		if x%y == 0 {
			return false, y
		}
	}
	return true, 0
}

func main() {

	var result int = 0
	var theNumber int = 600851475143

	for a := theNumber; a > 2; a = a - 2 {
		current, smaller := IsFactor(a)
		fmt.Println(a, " isPrime: ", current, " divided by ", smaller)
		if current {
			result = a
			break
		}
	}

	fmt.Println("The largest prime factor of ", theNumber, " is ", result)
}
