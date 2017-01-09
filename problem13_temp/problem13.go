package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type LargeNumberStr [100]string

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
	//return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
	//return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

func LargeNumberAdd(number1 string, number2 string) string {
	fmt.Println(number1, "  ", number2)
	r1 := []rune(number1)
	r2 := []rune(number2)
	r1_length := len(r1)

	result := make([]int, r1_length+1)
	r := 0
	for i := r1_length - 1; i >= 0; i-- {
		k, _ := strconv.Atoi(string(r1[i]))
		m, _ := strconv.Atoi(string(r2[i]))

		t := k + m + r
		if t >= 10 {
			t = t - 10
			r = 1
		} else {
			r = 0
		}
		result[i+1] = t

		if i == 0 {
			result[i] = r
		}
	}

	if result[0] == 0 {
		result = result[1:]
	}
	return strings.Trim(strings.Replace(fmt.Sprint(result), " ", "", -1), "[]")
}

func main() {
	//set up the data
	var data LargeNumberStr
	f, _ := os.Open("problem13data.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	result := ""
	for scanner.Scan() {
		data[i] = scanner.Text()
		toAdd := data[i]

		if len(result) > len(data[i]) {
			k := len(result) - len(data[i])
			for j := 0; j < k; j++ {
				toAdd = "0" + toAdd
			}
		}

		if i > 0 {
			result = LargeNumberAdd(result, toAdd)
		} else if i == 0 {
			result = toAdd
		}
		i++
	}

	fmt.Println(result)

}
