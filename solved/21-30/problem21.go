package main

import "fmt"
import "math"

func SumOfDivisors(n int) int {
	sum := 0
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			//get the number
			if i == n/i {
				sum += i
			} else {
				sum += i + (n / i)
			}
		}
	}

	return sum - n
}

func main() {
	sum := 0
	for i := 1; i < 10000; i++ {
		a := SumOfDivisors(i)
		b := SumOfDivisors(a)
		if i == b && i != a {
			sum += i
		}
	}

	fmt.Println(sum)

}
