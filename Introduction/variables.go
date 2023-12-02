package main

import "fmt"

func main() {
	//1: print hello world
	fmt.Println("Hello World")

	//2: create a variable with type inference
	z := 10
	fmt.Println(z)

	//3: create a constant
	const a int = 10
	fmt.Println("value of constant a  = ", a )

	//4: create a constant with type inference
	const b = 20
	fmt.Println("value of constant b  = ", b )

	//5: create an array and set values	
	var arr [5]int
	arr[0] = 10 
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("value of array arr  = ", arr )

	//6: create a set of variables
	var (
		c = 10
		d = 20
		e = 30
	)
	fmt.Println("value of variables c, d, e  = ", c, d, e )

	//7: 


}