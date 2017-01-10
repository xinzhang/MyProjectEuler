package main

import "fmt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	a, b, c := 1, 1, 1
	answerFound := false

	for a = 1; a < 500; a++ {
		for b = 1; b < ((1000 - a) / 2); b++ {
			c = 1000 - a - b
			if a*a+b*b == c*c {
				answerFound = true
				break
			}
		}
		if answerFound {
			break
		}
	}

	fmt.Println("There exists exactly one Pythagorean triplet for which a + b + c = 1000 ", a, "|", b, "|", c)
	fmt.Println("The product of a,b,c is", a*b*c)
}
