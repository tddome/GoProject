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

// actually fill in stuff, especially for order
func userAuthenticated(id int) {
	var i int
	var uAuth = 0
	i = GetIndexOfUser(id)

	for uAuth < 1 {
		if i == -1 {
			fmt.Println("Not a valid user ID.")
			uAuth = 1
		} else {
			fmt.Println("\nWelcome, user! What do you want to do?")
			fmt.Println("1. View user information")
			fmt.Println("2. Update email")
			fmt.Println("3. Update password")
			fmt.Println("4. Place order")
			fmt.Println("5. Access previous orders")
			fmt.Println("6. Delete account")
			fmt.Println("7. Sign out")
			var text string
			fmt.Scan(&text)

			switch text {
			case "1":
				fmt.Println("Your user information:\n")
				UserToString(id)
			case "2":
				fmt.Println("Enter new email address: testEmail@test.email")
				UserUpdateEmail(id, "testEmail@test.email")
				fmt.Println("Email updated.\n")
			case "3":
				fmt.Println("Enter new password: badPassword")
				UserUpdatePassword(id, "badPassword")
				fmt.Println("Password updated.\n")
			case "4":
				fmt.Println("Starting order...\n")
				NewOrder()
				fmt.Println("Adding product:\n")
				fmt.Println("Product ID: ")
				fmt.Println("Quantity: ")
				OrderAddProduct(1, 1)
				fmt.Println("Adding product:\n")
				fmt.Println("Product ID: ")
				fmt.Println("Quantity: ")
				OrderAddProduct(1, 1)
				fmt.Println("Remove product:\n")
				OrderRemoveProduct(1, 1)
				fmt.Println("Update quantity:\n")
				fmt.Println("Product:")
				OrderQuantityUpdate(1, 1)
				fmt.Println("Current order:\n")
				PrintOrder()
				fmt.Println("Finalizing order\n")
				fmt.Println("Discount code: code")
				FinalizeOrder(id, "discountcode")
				fmt.Println("Order complete.")
			case "5":
				fmt.Println("Your previous orders:\n")
				AccessPastOrders(id)
			case "6":
				fmt.Println("Deleting your account...")
				DeleteUser(id)
				fmt.Println("Account deleted.  Goodbye!\n")
				uAuth = 1
			case "7":
				fmt.Println("Signing out...\n")
				uAuth = 1
			default:
				fmt.Println("Incorrect input. Please try again.")
			}
		}
	}
}
