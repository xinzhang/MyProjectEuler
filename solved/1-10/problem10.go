package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsPrime(x int) bool {
	for y := 2; y <= x/2; y++ {
		if x%y == 0 {
			return false
		}
	}
	return true
}

func SumOfPrime(n int) int64 {

	arr := make([]bool, n)
	arr[0], arr[1] = true, true

	sum, prime := int64(5), 3
	var k int

	for {
		for k = 2 * prime; k < n; k += prime {
			arr[k] = true
		}
		for k = prime + 2; k < n && arr[k]; k += 2 {
		}
		if k < n {
			prime = k
			sum += int64(k)
		} else {
			break
		}
	}

	return sum

}

func main() {
	s := SumOfPrime(10)
	fmt.Println("The sum of the primes below 10 ", s)

	fmt.Println("The sum of the primes below two million ", SumOfPrime(2000000))
}
