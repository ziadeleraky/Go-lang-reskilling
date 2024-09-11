package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func writeBalanceToFile(balance float64) {
	os.WriteFile("balance.txt", []byte(fmt.Sprintf("%f", balance)), 0644);
}

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("balance.txt");

	if err != nil {
		return 0.0, errors.New("Error reading balance file");
	}

	balanceText := string(data);
	balance, err := strconv.ParseFloat(balanceText, 64);

	if err != nil {
		return 0.0, errors.New("Error parsing balance");
	}

	return balance, nil;
}

func main() {
	var balance, err = getBalanceFromFile();

	if err != nil {
		fmt.Println("Error reading balance file");
		fmt.Println(err);
		return;
	}

	fmt.Println("Welcome to Go bank!");
	fmt.Println("What do you want to do?");
	fmt.Println("1. Deposit");
	fmt.Println("2. Withdraw");
	fmt.Println("3. Check Balance");
	fmt.Println("4. Exit");

	var choice int;
	fmt.Print("Enter your choice: ");
	fmt.Scan(&choice);

	if choice == 1 {
		var deposit float64;
		fmt.Print("Enter the amount you want to deposit: ");
		fmt.Scan(&deposit);
		balance += deposit;
		fmt.Println("Deposited successfully! Your new balance is: ", balance);
	} else if choice == 2 {
		var withdraw float64;
		fmt.Print("Enter the amount you want to withdraw: ");
		fmt.Scan(&withdraw);
		if withdraw > balance {
			fmt.Println("Insufficient balance!");
		} else {
			balance -= withdraw;
			fmt.Println("Withdrawn successfully! Your new balance is: ", balance);
		}
	} else if choice == 3 {
		fmt.Println("Your balance is: ", balance);
	} else if choice == 4 {
		fmt.Println("Thank you for using Go bank!");
	} else {
		fmt.Println("Invalid choice!");
		return;
	}

	writeBalanceToFile(balance);
}