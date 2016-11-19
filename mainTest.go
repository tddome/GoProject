// CS408_Go_Project project main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	JamesBank := payAccount{"test", 1, 1}
	James := user{1, "password", "email", JamesBank}
	//DaveBank := payAccount{"test", 1, 1}
	//Dave := user{1, "password", "email", DaveBank}
	AddUserToDatabase(James)
	//AddUserToDatabase(Dave)
	fmt.Println(ReturnUserListLength())

	//something()

	//testProduct := Product{"pie", 11223, "food", 3.14}
	//ProductToString(testProduct)

	//fmt.Println(James.id)
}
