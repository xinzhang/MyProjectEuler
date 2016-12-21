package main

import "fmt"

func main() {
	var sum int = 2

	var firstVal int = 1
	var secondVal int = 2
	var currentVal int = 0

	for {
		//calcuate
		currentVal = (firstVal + secondVal)
		if currentVal > 4000000 {
			fmt.Println("The current val is ", currentVal)
			break
		}

		if currentVal%2 == 0 {
			sum += currentVal
		}

		//swap
		firstVal = secondVal
		secondVal = currentVal
		currentVal = 0

	}

	fmt.Println("For fibonacci sequence do not exceed 4 million, sum of all the even-value terms = ", sum)
}
