package main

import (
	"fmt"
)

// struct to hold user data
type userData struct {
	firstName string
	lastName string
	age string
}

func main() {
	// get user data
	firstName := getUserData("Enter your first name: ");
	lastName := getUserData("Enter your last name: ");
	age := getUserData("Enter your age: ");

	// create user data struct
	userData := userData{firstName, lastName, age};

	// output user data
	outputUserData(userData.firstName, userData.lastName, userData.age);

}

func outputUserData(firstName string, lastName string, age string) {
	fmt.Println("First Name: " + firstName);
	fmt.Println("Last Name: " + lastName);
	fmt.Println("Age: ", age);
}

func getUserData(promptText string) string {
	fmt.Print(promptText);
	var input string;
	fmt.Scan(&input);
	return input;
}