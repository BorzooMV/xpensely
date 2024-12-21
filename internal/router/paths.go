package router

import "fmt"

var v1Users = "api/v1/users"
var v1Expenses = "api/v1/expenses"

var Paths = map[string]string{
	"users":                v1Users,
	"userWithId":           fmt.Sprintf("%v/:id", v1Users),
	"expenses":             v1Expenses,
	"expensesOfSingleUser": fmt.Sprintf("%v/:id", v1Expenses),
}
