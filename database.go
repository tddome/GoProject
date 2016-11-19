// CS408_Go_Project project main.go

package main

var userList []user
var productList []product

func AddUserToDatabase(u user) {
	userList = append(userList, u)
}

//ProductDelete - Delete a Product from the Product list using index
//func ProductDelete(p []product, i int) []product {
//	return append(p[:i], p[i+1:]...)
//}

func ReturnUserListLength() int {
	return len(userList)
}

func AddProductToDatabase(p product) {
	productList = append(productList, p)
}

//ProductDelete - Delete a Product from the Product list using index
//func ProductDelete(p []product, i int) []product {
//	return append(p[:i], p[i+1:]...)
//}

func ReturnProductListLength() int {
	return len(productList)
}
