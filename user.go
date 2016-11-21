// CS408_Go_Project project main.go
package main

import (
	"fmt"
	"time"
)

type payAccount struct {
	bankName      string
	accountNumber int
	routingNumber int
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
	bAccNum int, bRoutNum int) {
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
	fmt.Println("\n")
}

func AccessPastOrders(id int) {
	for _, ord := range orderHistory {
		if ord.oUserID == id {
			OrderToString(ord.oBillID)
		}
	}

}

// actually fill in stuff, especially for order
func UserAuthenticated(id int) {
	var i int
	var uAuth = 0
	i = GetIndexOfUser(id)

	for uAuth < 1 {
		if i == -1 {
			fmt.Println("Not a valid user ID.")
			uAuth = 1
		} else {
			fmt.Println("Welcome, user!\n")
			fmt.Println("What do you want to do?")
			fmt.Println("1. View user information")
			fmt.Println("2. Update email")
			fmt.Println("3. Update password")
			fmt.Println("4. Place order")
			fmt.Println("5. Access previous orders")
			fmt.Println("6. Delete account")
			fmt.Println("7. Sign out\n")
			var text string
			fmt.Printf("-> ")
			fmt.Scan(&text)

			switch text {
			case "1":
				fmt.Println("\nYour user information:\n")
				UserToString(id)
			case "2":
				fmt.Println("\nEnter new email address: testEmail@test.email")
				fmt.Println("Updating email...\n")
				time.Sleep(2 * time.Second)
				UserUpdateEmail(id, "testEmail@test.email")
				fmt.Println("Email updated.\n")
			case "3":
				fmt.Println("\nEnter new password: badPassword")
				fmt.Println("Updating password...\n")
				time.Sleep(2 * time.Second)
				UserUpdatePassword(id, "badPassword")
				fmt.Println("Password updated.\n")
			case "4":
				fmt.Println("\nStarting order...\n")
				NewOrder()
				fmt.Println("Current products offered:\n")
				CheckProductList()
				var fOrd int = 0

				for fOrd < 1 {
					fmt.Println("What do you want to do?")
					fmt.Println("1. Add product")
					fmt.Println("2. Remove product")
					fmt.Println("3. Update quantity")
					fmt.Println("4. View current order")
					fmt.Println("5. Finish order")
					fmt.Println("6. Cancel order\n")
					var text2 string
					fmt.Printf("-> ")
					fmt.Scan(&text2)

					switch text2 {
					case "1":
						fmt.Println("\nProduct ID: 99")
						fmt.Println("Quantity: 1")
						fmt.Println("Adding product...\n")
						time.Sleep(2 * time.Second)
						OrderAddProduct(99, 1)
						fmt.Println("Product added.\n")

						// Extra products for the sake of demo
						OrderAddProduct(203, 7)
					case "2":
						fmt.Println("\nProduct ID: 99")
						fmt.Println("Removing product...\n")
						time.Sleep(2 * time.Second)
						OrderRemoveProduct(99)
						fmt.Println("Product removed.\n")
					case "3":
						fmt.Println("\nProduct ID: 203")
						fmt.Println("New quantity: 2")
						fmt.Println("Updating...\n")
						time.Sleep(2 * time.Second)
						OrderQuantityUpdate(203, 2)
						fmt.Println("Quantity updated.\n")
					case "4":
						fmt.Println("\nAccessing order...\n")
						time.Sleep(2 * time.Second)
						fmt.Println("Your current order:\n")
						PrintOrder()
					case "5":
						fmt.Println("\nDiscount code: SWEET30")
						fmt.Println("Finalizing order...\n")
						time.Sleep(2 * time.Second)
						FinalizeOrder(id, "SWEET30")
						fmt.Println("Order complete.")
						fOrd = 1
					case "6":
						fmt.Println("\nOrder cancelled.\n")
						fOrd = 1
					default:
						fmt.Println("\nIncorrect input. Please try again.")
					}
				}
			case "5":
				fmt.Println("\nAccessing orders...\n")
				time.Sleep(2 * time.Second)
				fmt.Println("Your previous orders:\n")
				AccessPastOrders(id)
			case "6":
				fmt.Println("\nDeleting your account...\n")
				time.Sleep(2 * time.Second)
				DeleteUser(id)
				fmt.Println("Account deleted.  Goodbye!\n")
				uAuth = 1
			case "7":
				fmt.Println("\nSigning out...\n")
				uAuth = 1
			default:
				fmt.Println("\nIncorrect input. Please try again.")
			}
		}
	}
}
