package main

import (
	"github.com/yousafgill/letsgo/pkg/controlflow"
	"github.com/yousafgill/letsgo/pkg/function"
	"github.com/yousafgill/letsgo/pkg/operator"
	"github.com/yousafgill/letsgo/pkg/variable"
)

func main() {

	//1: print hello world
	println("Hello World")

	//Display PrintVariables here
	operator.PrintOperators()

	//Display PrintVariables here
	variable.PrintVariables()

	//Display PrintControlFlow here
	controlflow.PrintControlFlow()
	
	//Display PrintFunctions here
	function.PrintFunctions()


}