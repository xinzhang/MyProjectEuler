package main

import "fmt"

func main() {
	var sum int = 0

	for x := 3; x < 1000; x++ {
		if x%3 == 0 || x%5 == 0 {
			sum += x
		}
	}

	fmt.Println("sum of all the multipes of 3 or 5 below 1000 = ", sum)
}
