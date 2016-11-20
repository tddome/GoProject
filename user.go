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
	var count int = 0

	for count < len(userList) {
		if userList[count].uID == id {
			fmt.Println("That user id is taken.  Please choose a new ID.")
			return
		} else {
			count++
		}
	}

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

func AccessPastOrders(id int) {
	for _, ord := range orderHistory {
		if ord.oUserID == id {
			OrderToString(ord.oBillID)
		}
	}

}
