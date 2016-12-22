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

type Grid [20][20]int

func main() {
	//set up the data
	var data Grid
	f, _ := os.Open("problem13data.txt")
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

	fmt.Println(result)

}
