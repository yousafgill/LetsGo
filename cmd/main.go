package main

import (
	"github.com/yousafgill/letsgo/pkg/controlflow"
	"github.com/yousafgill/letsgo/pkg/operators"
	"github.com/yousafgill/letsgo/pkg/variables"
)

func main() {

	//1: print hello world
	println("Hello World")

	//Display PrintVariables here
	operators.PrintOperators()

	//Display PrintVariables here
	variables.PrintVariables()

	//Display PrintControlFlow here
	controlflow.PrintControlFlow()
	


}