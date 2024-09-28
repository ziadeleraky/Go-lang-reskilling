package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	var count int

	for i := 0; i < n; i++ {
		var a, b, c int
		fmt.Scan(&a, &b, &c)
		if a+b+c >= 2 {
			count++
		}
	}

	fmt.Println(count)
}