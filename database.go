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
*	11/20/2016
*
* Function:
*	Database to store all users, products, discounts, and
*	orders created.  Includes addition, deletion, length,
*	and updating methods for each array.
 */

package main

// Arrays for storing information in the database
var userList []user
var productList []product
var discountList []discount
var orderHistory []order

// AddUserToDatabase - Adds a new user to the userList
func AddUserToDatabase(u user) {
	userList = append(userList, u)
}

// DeleteUserFromDatabase - Deletes a user from the userList
// using index
func DeleteUserFromDatabase(i int) {
	userList = append(userList[:i], userList[i+1:]...)
}

// GetUserLength - Returns an integer value of the length
// of userList
func GetUserLength() int {
	return len(userList)
}

// AddProductToDatabase - Adds a new product to the productList
func AddProductToDatabase(p product) {
	productList = append(productList, p)
}

// ProductDelete - Deletes a product from the productList
// using index
func DeleteProductFromDatabase(i int) {
	productList = append(productList[:i], productList[i+1:]...)
}

// GetProductLength - Returns an integer value of the length
// of productList
func GetProductLength() int {
	return len(productList)
}

// AddDiscountToDatabase - Adds a new discount to the discountList
func AddDiscountToDatabase(d discount) {
	discountList = append(discountList, d)
}

// DeleteDiscountFromDatabase - Deletes a discount from the
// discountList using index
func DeleteDiscountFromDatabase(i int) {
	discountList = append(discountList[:i], discountList[i+1:]...)
}

// GetDiscountLength - Returns an integer value of the length
// of discountList
func GetDiscountLength() int {
	return len(discountList)
}

// AddOrderToDatabase - Adds a new order to the orderHistory
func AddOrderToDatabase(o order) {
	orderHistory = append(orderHistory, o)
}

// GetOrderHistoryLength - Returns an integer value of the length
// of orderHistory
func GetOrderHistoryLength() int {
	return len(orderHistory)
}
