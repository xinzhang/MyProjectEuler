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

func main() {
	data, err := ioutil.ReadFile("p022_names.txt")
	check(err)

	tmp
	cnt := 0
	for i := 0; i < len(data); i++ {
		s := string(data[i])
		if s == "," {
			cnt++
		}
	}

	fmt.Println(cnt)
}
