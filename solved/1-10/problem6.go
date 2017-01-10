package main

import "fmt"

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func SumOfSquare(y int) int {
	sum := 0
	for i := 1; i <= y; i++ {
		sum += (i * i)
	}
	return sum
}

func SquareOfSum(y int) int {
	sum := 0
	for i := 1; i <= y; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	fmt.Println(SumOfSquare(10))
	fmt.Println(SquareOfSum(10))
	fmt.Println("ind the difference between the sum of the squares of the first 10 natural numbers and the square of the sum", SumOfSquare(10)-SquareOfSum(10))

	fmt.Println(SumOfSquare(100))
	fmt.Println(SquareOfSum(100))
	fmt.Println("ind the difference between the sum of the squares of the first 10 natural numbers and the square of the sum", SumOfSquare(100)-SquareOfSum(100))
}
