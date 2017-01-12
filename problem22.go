package main

import (
	"fmt"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SortStringArray(current []string, val string) []string {
	result := make([]string, len(current)+1)

}

func main() {
	data, err := ioutil.ReadFile("p022_names.txt")
	check(err)

	tmp := make([]byte, 0)
	cnt := 0
	for i := 0; i < len(data); i++ {
		s := string(data[i])
		if s == "," {
			//stopped and show the tmp
			fmt.Println(string(tmp))
			tmp = make([]byte, 0)
			cnt++
		} else if s == "\"" {
			continue
		} else {
			tmp = append(tmp, data[i])
		}

	}

	fmt.Println(cnt)
}
