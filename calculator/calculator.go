package main

import (
	"fmt"

	"example.com/calculator/ops"
)

func main() {
	var a, b int;
	fmt.Println("Enter two numbers: ");
	fmt.Scan(&a, &b);

	fmt.Println("Addition: ", ops.Add(a, b));
	fmt.Println("Subtraction: ", ops.Subtract(a, b));
	fmt.Println("Multiplication: ", ops.Multiply(a, b));
	fmt.Println("Division: ", ops.Divide(a, b));
}