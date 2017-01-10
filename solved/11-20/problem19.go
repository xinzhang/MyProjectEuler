package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	for i := 1901; i < 2001; i++ {
		for j := 1; j < 13; j++ {
			start := time.Date(i, time.Month(j), 1, 0, 0, 0, 0, time.UTC)
			if start.Weekday() == time.Sunday {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
