// GoProject

//Contributors:
// Shaylyn Wetts
// Troy Dome

// References:
// https://www.reddit.com/r/golang/comments/2uko6l/algorithms_and_data_structures_implemented_in_go/
// https://www.goinggo.net/2013/09/iterating-over-slices-in-go.html
// http://stackoverflow.com/questions/7782411/is-there-a-foreach-loop-in-go

// Project Outline (Might only work for Troy?):
// https://docs.google.com/document/d/1ocq3TYRU3JASzWDYsycaDi9CGuoG2MVsmT2LDAnvnaE/edit?ts=580ed160

package main

//	"fmt"

func main() {
	// create manager; show manager options
	ManagerAuthenticated()

	// DEMO: input 1
	// add product (display product info for adding)
	// add more products (do not display info)

	// DEMO: input 3
	// check product list

	// DEMO: input 2
	// remove product

	// DEMO: input 3
	// check product list again

	// DEMO: input 4
	// add discount (display)
	// add other discounts

	// DEMO: input 6
	// check discount list

	// DEMO: input 5
	// remove discount

	// DEMO: input 6
	// check discount list to show changes

	// DEMO: input 8
	// exit manager

	// create a bunch of users
	CreateUser(1432, "danielEmail@email.email", "myPassword", "My Bank",
		1122334455, 6677889900)
	CreateUser(8945, "stupidguy@smart.gov", "BigBoy49", "Bill Busters",
		348235, 239145)
	CreateUser(3852, "badjob@goodjob.info", "iLoveMyJob22", "Bill Crushers",
		219445, 459512)

	// login to user; show user options
	UserAuthenticated(1432)

	// DEMO: input 1
	// output user information

	// DEMO: input 2
	// change email

	// DEMO: input 3
	// change password

	// DEMO: input 1
	// output user information to show changes

	// DEMO: input 4
	// create order

	// DEMO: input 1
	// add products to order

	// DEMO: input 4
	// show products in order with quantity

	// DEMO: input 2
	// delete product from order

	// DEMO: input 3
	// update product quantity

	// DEMO: input 4
	// display current order

	// DEMO: input 5
	// finalize order

	// DEMO: input 5
	// show previous orders

	// DEMO: input 6
	// delete account
}
