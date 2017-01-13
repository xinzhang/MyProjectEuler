package main

import "math"
import "container/list"

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

var writeable []bool
var limit int
var abunList *list.List

func checkWritables() {
	for i := abunList.Front(); i != nil; i = i.Next() {
		for j := i; j != nil; j = j.Next() {
			if sum := i.Value.(int) + j.Value.(int); sum < limit {
				writeable[sum] = true
			}
		}
	}
}

func main() {
	abunList = list.New()
	limit = 28124
	writeable = make([]bool, limit)

	for i := 0; i < limit; i++ {
		writeable[i] = false
		if i < SumOfDivisors(i) {
			abunList.PushBack(i)
		}
	}

	checkWritables()

	sum := 0
	for i := 1; i < limit; i++ {
		if !writeable[i] {
			sum += i
		}
	}

	println(sum)
}
