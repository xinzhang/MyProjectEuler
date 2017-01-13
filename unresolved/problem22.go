package main

import (
	"fmt"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SortStringArray(current []string, val string) []string {
	//fmt.Println(current, val)

	result := make([]string, len(current)+1)
	if len(current) == 0 {
		result[0] = val
		fmt.Println(result)
		return result
	}

	i := 0
	for i = 0; i < len(current); i++ {
		//fmt.Println(current[i], val)

		if strings.Compare(current[i], val) == -1 {
			continue
		} else if strings.Compare(current[i], val) <= 0 {
			if i == 0 {
				result = append()
			} else if i == 1 {
				result = append(current[0:i-2], val)
				result = append(result, current[i-1:]...)
			} else {
				result = append(current, val)
				result = append(result, current[i:len(current)-1]...)
			}

			fmt.Println("dsfsdafs", result)
			return result
		}
	}

	if i == len(current)-1 {
		fmt.Println("current", current)
		result = append(current, val)
		fmt.Println("result", result)
	}

	//fmt.Println(result)
	return result
}

func main() {

	//data, err := ioutil.ReadFile("p022_names.txt")
	//check(err)
	sortedArray := make([]string, 0)
	sortedArray = SortStringArray(sortedArray, "Tuan")
	fmt.Println(sortedArray)
	sortedArray = SortStringArray(sortedArray, "Crange")
	fmt.Println(sortedArray)
	sortedArray = SortStringArray(sortedArray, "Pipe")
	fmt.Println(sortedArray)

	/*
		tmp := make([]byte, 0)
		cnt := 0
		for i := 0; i < len(data); i++ {
			s := string(data[i])
			if s == "," {
				//stopped and show the tmp
				//fmt.Println(string(tmp))
				sortedArray = SortStringArray(sortedArray, string(tmp))
				fmt.Println(string(tmp))
				tmp = make([]byte, 0)
				cnt++
			} else if s == "\"" {
				continue
			} else {
				tmp = append(tmp, data[i])
			}

			if i == 7 {
				break
			}

		}

		//fmt.Println(cnt)
		//fmt.Println(sortedArray)
	*/

}
