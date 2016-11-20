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
	uID          int
	uEmail       string
	uPassword    string
	uBankAccount payAccount
}

func GetIndexOfUser(id int) int {
	for i, a := range userList {
		if a.uID == id {
			return i
		}
	}
	return -1
}

func CreateUser(id int, email string, pass string, bName string,
	bAccNum float64, bRoutNum float64) {

	pANew := payAccount{
		bankName:      bName,
		accountNumber: bAccNum,
		routingNumber: bRoutNum,
	}
	uNew := user{
		uID:          id,
		uEmail:       email,
		uPassword:    pass,
		uBankAccount: pANew,
	}
	AddUserToDatabase(uNew)
}

func DeleteUser(id int) {
	var i int
	i = GetIndexOfUser(id)
	DeleteUserFromDatabase(i)
}

func UserUpdateEmail(id int, n string) {
	var i int
	i = GetIndexOfUser(id)
	userList[i].uEmail = n
}

func UserUpdatePassword(id int, n string) {
	var i int
	i = GetIndexOfUser(id)
	userList[i].uPassword = n
}

func UserToString(id int) {
	var i int
	i = GetIndexOfUser(id)

	fmt.Println("User Id:              ", userList[i].uID)
	fmt.Println("User Email:         ", userList[i].uEmail)
	fmt.Println("User Password:    ", userList[i].uPassword)
	fmt.Println("User Bank Name: ", userList[i].uBankAccount.bankName)
	fmt.Println("User Account #:  ", userList[i].uBankAccount.accountNumber)
	fmt.Println("User Routing #:  ", userList[i].uBankAccount.routingNumber)
}

func PlaceOrder() {
	// references the Order thing
}

func AccessPastOrders() {
	// lists all past order info made by this user
}
