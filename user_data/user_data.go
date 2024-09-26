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

// constructor
func newUser(firstName string, lastName string, age string) *userData {
	return &userData{firstName, lastName, age};
}

// method to clear user data
func (user *userData) clearUserData() {
	user.firstName = "";
	user.lastName = "";
	user.age = "";
}

func main() {
	// create user data struct
	user := newUser("Ziad", "ElEraky", "26");
	// // get user data
	// firstName := getUserData("Enter your first name: ");
	// lastName := getUserData("Enter your last name: ");
	// age := getUserData("Enter your age: ");

	// // create user data struct
	// user := userData{firstName, lastName, age};

	user.outputUserData();
	user.clearUserData();
	user.outputUserData();
}

func (user userData) outputUserData() {
	fmt.Println("First Name: " + user.firstName);
	fmt.Println("Last Name: " + user.lastName);
	fmt.Println("Age: ", user.age);
}

func getUserData(promptText string) string {
	fmt.Print(promptText);
	var input string;
	fmt.Scan(&input);
	return input;
}