package main

import "fmt"

func main() {
	n := 1000000
	var result []int

	theNumber := 0
	max := 1

	result = append(result, 1)

	for i := 1; i < n; i++ {
		currentVal := i
		step := 0
		for {
			if currentVal == 1 {
				break
			}

			if currentVal%2 == 0 {
				currentVal = currentVal / 2
			} else {
				currentVal = currentVal*3 + 1
			}

			step++
			if step > max {
				theNumber = i
				max = step
			}
		}
	}

	fmt.Println(max)
	fmt.Println(theNumber)
}
