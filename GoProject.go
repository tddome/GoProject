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
*	Main method for testing and simulating the functionality
*	of the project.  Includes comments on what to do and
*	what the results will be for in class presentation demo.
 */

package main

import (
	"fmt"
)

func main() {

	// Stock products for testing
	CreateProduct("BEAR", 99, "ANIMAL", 1000000)
	CreateProduct("PAINT BRUSH", 406, "ART SUPPLIES", 15)
	CreateProduct("PHILLY CHEESE STEAK", 101, "FOOD", 7)
	CreateProduct("SANDALS", 51, "CLOTHING", 20)
	CreateProduct("DOG TOY DELUXE", 203, "TOY", 30)
	CreateProduct("CLASSIC TANK TOP", 309, "CLOTHING", 90)

	// Stock users for testing
	CreateUser(1432, "danielEmail@email.email", "myPassword", "My Bank",
		1122334455, 6677889900)
	CreateUser(8945, "stupidguy@smart.gov", "BigBoy49", "Bill Busters",
		348235, 239145)
	CreateUser(3852, "badjob@goodjob.info", "iLoveMyJob22", "Bill Crushers",
		219445, 459512)

	// Stock discounts for testing
	CreateDiscount("WOWMAN", 70)
	CreateDiscount("BEARSPECIAL", 40)
	CreateDiscount("SWEET30", 30)
	CreateDiscount("CHRISTMASJOY", 50)

	// Switch statement for choosing between logging in as a
	// manager, creating a new user, and logging in as a user
	var oAuth = 0
	for oAuth < 1 {
		fmt.Println("Please choose an option:")
		fmt.Println("1. Sign in as manager")
		fmt.Println("2. Create user account")
		fmt.Println("3. Sign in as user")
		fmt.Println("4. Exit")
		var text string
		fmt.Printf("\n-> ")
		fmt.Scan(&text)

		switch text {
		case "1":
			ManagerAuthenticated()
		case "2":
			var inID int
			var inEmail string
			var inPass string
			var inBank string
			var inAccount int
			var inRouting int

			fmt.Println("\nCreate new user")
			fmt.Printf("\nEnter ID (four digits): ")
			fmt.Scan(&inID)
			fmt.Printf("Enter email: ")
			fmt.Scan(&inEmail)
			fmt.Printf("Enter password: ")
			fmt.Scan(&inPass)
			fmt.Printf("Enter bank name: ")
			fmt.Scan(&inBank)
			fmt.Printf("Enter account number: ")
			fmt.Scan(&inAccount)
			fmt.Printf("Enter routing number: ")
			fmt.Scan(&inRouting)
			CreateUser(inID, inEmail, inPass, inBank, inAccount, inRouting)
			fmt.Println("\nUser created.")
		case "3":
			fmt.Printf("Please enter a user id: ")
			var signInID int
			fmt.Scan(&signInID)
			UserAuthenticated(signInID)
		case "4":
			fmt.Println("Goodbye.")
			oAuth = 1
		default:
			fmt.Println("Incorrect input. Please try again.")
		}
	}

}
