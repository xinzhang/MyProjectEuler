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

type LargeNumberStr [50]string

func LargeNumberAdd(number1 string, number2 string) string {
	fmt.Println(number1, "  ", number2)
	r1 := []rune(number1)
	r2 := []rune(number2)

	var result [10]int
	r:=0;
	for i:=9; i<0; i-- {
			k,_ := strconv.Atoi(r1[i])
			m,_ := strconv.Atoi(r2[i])
			t := k+m+r
			r=t/10
			result[i] = t
	}

 	return string(result)
}

func main() {
	//set up the data
	var data LargeNumberStr
	f, _ := os.Open("problem13data.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	result := ""
	for scanner.Scan() {
		LargeNumberStr[i] = scanner.Text()
		if (i>0){
			result = LargeNumberAdd(result, LargeNumberStr[i])
		} else if (i==0) {
			result = LargeNumberStr[i];
		}
		i++
	}

	fmt.Println(result)

}
