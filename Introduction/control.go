package letsgo

func control() {

	a := 10
	b := 20
	c := 15

	//1: if statement
	println("1: if statement")
	if a > b {
		println("1: a is greater than b ")
	}

	//2: if else statement
	println("2: if else statement")
	if a > b {
		println("1: a is greater than b")
	} else {
		println("2: b is greater than a")
	}

	//3: if else if statement
	println("3: if else if statement")
	if a > b {
		println("1: a is greater than b")
	} else if b > c {
		println("2: b is greater than c")
	} else {
		println("3: c is greater than a and b")
	}

}