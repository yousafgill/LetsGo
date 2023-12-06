package main

import "fmt"

func variables() {
	//1: print hello world
	fmt.Println("Hello World")

	//2: create a variable with type inference
	z := 10
	fmt.Println(z)

	//3: create a constant
	const a int = 10
	fmt.Println("value of constant a  = ", a)

	//4: create a constant with type inference
	const b = 20
	fmt.Println("value of constant b  = ", b)

	//5: create an array and set values
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("value of array arr  = ", arr)

	//6: create a set of variables
	var (
		c = 10
		d = 20
		e = 30
	)
	fmt.Println("value of variables c, d, e  = ", c, d, e)

	//7: Create a boolean variable
	var f bool = true
	fmt.Println("value of variable f  = ", f)

	//8: Create a float variable
	var g float32 = 10.5
	fmt.Println("value of variable g  = ", g)

	//9: create a complex128 variable
	var h complex128 = 10 + 5i
	fmt.Println("value of variable h  = ", h)

	//10: create a struct variable
	type Person struct {
		name string
		age  int
	}
	var i Person = Person{"John", 30}
	fmt.Println("value of variable i  = ", i)

	//11: create a pointer variable
	var j *int
	k := 10
	j = &k
	fmt.Println("value of variable j  = ", j)
	fmt.Println("value of variable j  = ", *j)
	k = 30
	fmt.Println("value of variable j  = ", j)
	fmt.Println("value of variable j  = ", *j)

	//12: create a channel variable
	var l chan int = make(chan int)
	fmt.Println("value of variable l  = ", l)

	//13: create a map variable
	var m map[string]int = make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	fmt.Println("value of variable m  = ", m)

	//14: create a slice variable
	var n []int = make([]int, 5)
	n[0] = 10
	n[1] = 20
	n[2] = 30
	n[3] = 40
	n[4] = 50
	fmt.Println("value of variable n  = ", n)

	//15: create a interface variable
	var o interface{} = "Hello, Yousaf Gill"
	fmt.Println("value of variable o  = ", o)
	o = 10
	fmt.Println("value of variable o  = ", o)
	o = 10.5
	fmt.Println("value of variable o  = ", o)
	o = true
	fmt.Println("value of variable o  = ", o)
	o = Person{"Yousaf Gill", 30}
	fmt.Println("value of variable o  = ", o)

	//16: create a function variable
	var p func(int, int) int = func(a int, b int) int {
		return a + b
	}
	fmt.Println("value of variable p  = ", p(10, 20))

	//17: create a function variable with type inference
	q := func(a int, b int) int {
		return a + b
	}
	fmt.Println("value of variable q  = ", q(100, 200))

	//18: create a constant with iota
	const (
		r = iota
		s = iota
		t = iota
	)
	fmt.Println("value of constant r, s, t  = ", r, s, t)

}