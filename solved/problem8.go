package main

import "fmt"
import "strconv"
import "io/ioutil"

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

func main() {
	data, err := ioutil.ReadFile("problem8data.txt")
	check(err)

	dataString := string(data)

	//step := 4
	step := 13
	length := len(dataString)
	fmt.Println("total length", length)

	r := []rune(dataString)

	var maxResult int64 = 1

	listResult := make([]int, step)

	for i := 0; i <= (length - step); i++ {
		var tempVal int64 = 1
		for j := 0; j < step; j++ {
			tempTotal, _ := strconv.Atoi(string(r[i+j]))
			if tempTotal == 0 {
				break
			}
			tempVal = tempVal * int64(tempTotal)
		}
		if tempVal > maxResult {
			maxResult = tempVal
			for j := 0; j < step; j++ {
				listResult[j], _ = strconv.Atoi(string(r[i+j]))
			}
			fmt.Println(listResult, maxResult)
		}
	}

	fmt.Println(maxResult)
	fmt.Println(listResult)

}
