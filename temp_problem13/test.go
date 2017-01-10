package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fileBuf, err := ioutil.ReadFile("problem13data.txt")
	if err != nil {
		panic(err.Error())
	}
	fileStr := string(fileBuf)
	lines := strings.Split(fileStr, "\n")
	result := ""
	carryout, digit := 0, 0
	for i := 49; i >= 0; i-- {
		sum := carryout
		for j := range lines {
			num, _ := strconv.Atoi(string(lines[j][i]))
			sum += num
		}
		digit = sum % 10
		carryout = sum / 10
		result = strconv.Itoa(digit) + result
	}
	result = strconv.Itoa(carryout) + result
	println(result[0:10])
}
