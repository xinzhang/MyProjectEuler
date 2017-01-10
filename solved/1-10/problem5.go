package main

import "fmt"

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func InSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func IndexInSlice(list []int, a int) int {
	for i, b := range list {
		if b == a {
			return i
		}
	}
	return -1
}

func MergeFactors(slice1 []int, slice2 []int) []int {
	result := make([]int, len(slice1))
	copy(result, slice1)

	slice2left := make([]int, 0)

	for _, num := range slice2 {
		if idx := IndexInSlice(result, num); idx >= 0 {
			result = append(result[:idx], result[idx+1:]...)
		} else {
			slice2left = append(slice2left, num)
		}
	}

	for _, num := range slice1 {
		slice2left = append(slice2left, num)
	}

	return slice2left
}

func GetAllFactors(n int) []int {
	result := make([]int, 0)

	num, t := n, 1

	for {
		t = GetNextReminder(num)
		result = append(result, t)
		num = num / t
		if 1 == num {
			break
		}
	}
	return result
}

func GetNextReminder(n int) int {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return i
		}
	}
	return n
}

func GetLeastCommonMultipe(n int) (int, []int) {
	start := make([]int, 0)
	for i := 2; i <= n; i++ {
		start = MergeFactors(start, GetAllFactors(i))
	}

	result := 1
	for _, v := range start {
		result = result * v
	}

	return result, start
}

func main() {
	/* test prime number
	fmt.Println(IsPrime(2))
	fmt.Println(IsPrime(3))
	fmt.Println(IsPrime(4))
	fmt.Println(IsPrime(5))
	fmt.Println(IsPrime(6))
	fmt.Println(IsPrime(7))
	*/

	/* test get all factors
	fmt.Println(GetAllFactors(10))
	fmt.Println(GetAllFactors(20))
	fmt.Println(GetAllFactors(30))
	*/

	/*
		v1 := []int{2, 3, 5}
		v2 := []int{2, 2, 5}
		fmt.Println(MergeFactors(v1, v2))
	*/

	result, list := GetLeastCommonMultipe(10)
	fmt.Println("the smallest positive number that is evenly divisible by all of the numbers from 1 to 10 ", result, list)

	result, list = GetLeastCommonMultipe(20)
	fmt.Println("the smallest positive number that is evenly divisible by all of the numbers from 1 to 20 ", result, list)
}
