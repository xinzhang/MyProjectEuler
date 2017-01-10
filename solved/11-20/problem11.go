package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Grid [20][20]int

func main() {
	//set up the data
	var data Grid
	f, _ := os.Open("problem11data.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		lineStr := scanner.Text()

		vals := strings.Fields(lineStr)
		for idx := range vals {
			data[i][idx], _ = strconv.Atoi(vals[idx])
		}
		i++
	}

	result := 0

	//from left to righ
	for i := 0; i < 20; i++ {
		for j := 0; j < 17; j++ {
			sum := data[i][j] * data[i][j+1] * data[i][j+2] * data[i][j+3]
			if sum > result {
				result = sum
			}
		}
	}

	//from left to righ
	for j := 0; j < 20; j++ {
		for i := 0; i < 17; i++ {
			sum := data[i][j] * data[i+1][j] * data[i+2][j] * data[i+3][j]
			if sum > result {
				result = sum
			}
		}
	}

	//cross
	for i := 0; i < 17; i++ {
		for j := 0; j < 17; j++ {
			sum := data[i][j] * data[i+1][j+1] * data[i+2][j+2] * data[i+3][j+3]
			if sum > result {
				result = sum
			}
		}
	}

	for i := 3; i < 20; i++ {
		for j := 3; j < 17; j++ {
			sum := data[i][j] * data[i-1][j+1] * data[i-2][j+2] * data[i-3][j+3]
			if sum > result {
				result = sum
			}
		}
	}

	fmt.Println(result)

}
