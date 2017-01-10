package main

import "fmt"
import "strconv"
import "math"

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func IsPalindrome(y int) bool {
	s := strconv.Itoa(y)
	i, _ := strconv.Atoi(StringReverse(s))
	return (i == y)
}

func FindLargetPalindrome(digit int) (result int, x int, y int) {

	min := int(math.Pow(10, float64(digit-1)))
	max := int(math.Pow(10, float64(digit)))

	result, x, y = 0, 0, 0
	for i := min; i < max; i++ {
		for j := min; j < max; j++ {
			if IsPalindrome(i * j) {
				if result <= i*j {
					result, x, y = i*j, i, j
				}
			}
		}
	}

	return result, x, y
}

func main() {
	fmt.Println(IsPalindrome(121))
	fmt.Println(IsPalindrome(1221))
	fmt.Println(IsPalindrome(12321))
	fmt.Println(IsPalindrome(1341))

	r, x, y := FindLargetPalindrome(2)
	fmt.Println("The palindrome made from the product of two 2-digit numbers ", r, x, y)
	r, x, y = FindLargetPalindrome(3)
	fmt.Println("The palindrome made from the product of two 3-digit numbers ", r, x, y)
	r, x, y = FindLargetPalindrome(4)
	fmt.Println("The palindrome made from the product of two 4-digit numbers ", r, x, y)
}
