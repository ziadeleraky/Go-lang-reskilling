package main

import "fmt"

func main() {
	var n int;
	fmt.Scan(&n);

	var arr []int;
	for i := 0; i < n; i++ {
		var number int;
		fmt.Scan(&number);
		arr = append(arr, number);
	}

	// max length of increasing subsequence
	maxLength := 1;
	// current length of increasing subsequence
	currentLength := 1;

	for i := 0; i < n-1; i++ {
		if (arr[i+1] > arr[i]) {
			currentLength++;
		} else {
			// if the current length of the subsequence is greater than the max length
			if (currentLength > maxLength) {
				maxLength = currentLength;
			}
			// reset current length because the subsequence is not increasing anymore
			currentLength = 1;
		}
	}

	// check if the current length of the subsequence is greater than the max length
	if (currentLength > maxLength) {
		maxLength = currentLength;
	}

	fmt.Println(maxLength)
}
