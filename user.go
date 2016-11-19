// CS408_Go_Project project main.go
package main

import (
	"fmt"
)

type payAccount struct {
	bankName      string
	accountNumber float64
	routingNumber float64
}

type user struct {
	id          float64
	email       string
	password    string
	bankAccount payAccount
}

func placeOrder() {
	// references the Order thing
}

func accessPastOrders() {
	// lists all past order info made by this user
}

func something() {
	Dave := user{id: 1145}
	Dave.email = "davesux@bademail.net"
	Dave.password = "iamatool69420"
	fmt.Println("We are testing Dave's shit! ")
	fmt.Printf("His id number is: ")
	fmt.Println(Dave.id)

	fmt.Printf("His email is: ")
	fmt.Println(Dave.email)

	fmt.Printf("His password is: ")
	fmt.Println(Dave.password)

	Dave.bankAccount.accountNumber = 111111111
	Dave.bankAccount.bankName = "Gimmie Your Money Co."
	Dave.bankAccount.routingNumber = 222222222

	fmt.Println("Dave's bank shit is...")
	fmt.Printf("Bank name: ")
	fmt.Printf(Dave.bankAccount.bankName)

	fmt.Printf("Account number: ")
	fmt.Println(Dave.bankAccount.accountNumber)

	fmt.Printf("Routing number: ")
	fmt.Println(Dave.bankAccount.routingNumber)
}
