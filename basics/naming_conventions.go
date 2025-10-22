package basics

import "fmt"

type Employee struct {
	FirstName string // PascalCase for struct fields
	LastName  string
	Age       int
}

func demonstrateNamingConventions() {
	// PascalCase
	// E.g. ClaculateArea, UserInfo, NewHttpRequest

	// Structs, interfaces, enums

	// snake_case
	// E.g. user_id, fist_name, http_request

	// UPPERCASE
	// Use case is Constants

	//mixedCase
	// E.g. calculateArea, userInfo, newHttpRequest

	const NAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ", employeeID)

}
