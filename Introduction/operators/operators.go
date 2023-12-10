package letsgo

func operators() {
	//1: arithmetic operators
	println("1: arithmetic operators")
	println("1 + 2 = ", 1+2)     //addition
	println("1 - 2 = ", 1-2)     //subtraction
	println("1 * 2 = ", 1*2)     //multiplication
	println("1 / 2 = ", 14/2)    //division
	println("1 % 2 = ", 1%2)     //modulus
	println("1 & 2 = ", 1&2)     //bitwise AND
	println("1 | 2 = ", 1|2)     //bitwise OR
	println("1 ^ 2 = ", 1^2)     //bitwise XOR
	println("1 << 2 = ", 1<<2)   //bitwise left shift
	println("1 >> 2 = ", 1>>2)   //bitwise right shift
	println("1 &^ 2 = ", 1&^2)   //bit clear (AND NOT)
	println("1 == 2 = ", 1 == 2) //equal to
	println("1 != 2 = ", 1 != 2) //not equal to
	println("1 < 2 = ", 1 < 2)   //less than
	println("1 <= 2 = ", 1 <= 2) //less than or equal to
	println("1 > 2 = ", 1 > 2)   //greater than
	println("1 >= 2 = ", 1 >= 2) //greater than or equal to

	var a bool = 10 == 410       //logical AND
	var b bool = 20 == 10        //logical OR
	println("1 && 2 = ", a && b) //logical AND

	println("1 || 2 = ", a || b) //logical OR
	println("1 &^ 2 = ", 1&^2)   //bit clear (AND NOT)
}