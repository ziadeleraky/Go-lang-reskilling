package main

import "fmt"

func main() {
	var n int;
	fmt.Scan(&n);

	for i := 0; i < n; i++ {
		var word string;
		fmt.Scan(&word);
		
		if (len(word) <= 10) {
			fmt.Println(word);
		} else {
			fmt.Println(word[0:1] + fmt.Sprint(len(word) - 2) + word[len(word) - 1:]);
		}
	}
}