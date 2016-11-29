/* GoProject
*
* Class:
*	CS 408
*
* Contributors:
* 	Troy Dome
* 	Shaylyn Wetts
*
* Last Updated:
*	11/28/2016
*
* Function:
*	Creates a user and pay account. Includes functions for
*	adding users, removing users, updating email and password,
*	and accessing past orders by this user.  Includes prompts
*	for user input to run these functions.
 */

package main

import (
	"fmt"
)

// PayAccount - Contains bank name, account number, and
// routing number for a user
type payAccount struct {
	bankName      string
	accountNumber int
	routingNumber int
}

// User - Contains ID, email, password, and bank account
// information
type user struct {
	uID          int
	uEmail       string
	uPassword    string
	uBankAccount payAccount
}

// GetIndexOfUser - Check if a User is contained by ID
// and returns the index where the ID is found, else return -1
func GetIndexOfUser(id int) int {
	for i, a := range userList {
		if a.uID == id {
			return i
		}
	}
	return -1
}

// CreateUser - Creates a user with a unique ID and adds the user
// to the database
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

// DeleteUser - Deletes a current user from the database
func DeleteUser(id int) {
	var i int
	i = GetIndexOfUser(id)
	DeleteUserFromDatabase(i)
}

// UserUpdateEmail - Updates the email of a user
func UserUpdateEmail(id int, n string) {
	var i int
	i = GetIndexOfUser(id)
	userList[i].uEmail = n
}

// UserUpdatePassword - Updates the password of a user
func UserUpdatePassword(id int, n string) {
	var i int
	i = GetIndexOfUser(id)
	userList[i].uPassword = n
}

// UserToString - Prints all information of a user
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

// AccessPastOrders - Prints all orders made by a specific user
func AccessPastOrders(id int) {
	for _, ord := range orderHistory {
		if ord.oUserID == id {
			OrderToString(ord.oBillID)
		}
	}

}

// UserAuthenticated - All functionality for a user
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
				var inEmail string

				fmt.Printf("\nEnter new email address: ")
				fmt.Scan(&inEmail)
				fmt.Println("Updating email...\n")
				UserUpdateEmail(id, inEmail)
				fmt.Println("Email updated.\n")
			case "3":
				var inPass string

				fmt.Printf("\nEnter new password: ")
				fmt.Scan(&inPass)
				fmt.Println("Updating password...\n")
				UserUpdatePassword(id, inPass)
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
						var inId int
						var inQuant int

						fmt.Printf("\nProduct ID: ")
						fmt.Scan(&inId)
						fmt.Printf("Quantity: ")
						fmt.Scan(&inQuant)
						fmt.Println("Adding product...\n")
						OrderAddProduct(inId, inQuant)
						fmt.Println("Product added.\n")
					case "2":
						var inId int

						fmt.Printf("\nProduct ID: ")
						fmt.Scan(&inId)
						fmt.Println("Removing product...\n")
						OrderRemoveProduct(inId)
						fmt.Println("Product removed.\n")
					case "3":
						var inId int
						var inQuant int

						fmt.Printf("\nProduct ID: ")
						fmt.Scan(&inId)
						fmt.Printf("New quantity: ")
						fmt.Scan(&inQuant)
						fmt.Printf("Updating...\n")
						OrderQuantityUpdate(inId, inQuant)
						fmt.Println("Quantity updated.\n")
					case "4":
						fmt.Println("\nAccessing order...\n")
						fmt.Println("Your current order:\n")
						PrintOrder()
					case "5":
						var inCode string

						fmt.Printf("\nDiscount code: ")
						fmt.Scan(&inCode)
						fmt.Println("Finalizing order...\n")
						FinalizeOrder(id, inCode)
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
				fmt.Println("Your previous orders:\n")
				AccessPastOrders(id)
			case "6":
				fmt.Println("\nDeleting your account...\n")
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
