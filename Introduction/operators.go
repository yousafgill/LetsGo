package main

func operators() {
	//1: arithmetic operators
	println("1: arithmetic operators")
	println("1 + 2 = ", 1+2)
	println("1 - 2 = ", 1-2)
	println("1 * 2 = ", 1*2)
	println("1 / 2 = ", 1/2)
	println("1 % 2 = ", 1%2)
	println("1 & 2 = ", 1&2)
	println("1 | 2 = ", 1|2)
	println("1 ^ 2 = ", 1^2)
	println("1 << 2 = ", 1<<2)
	println("1 >> 2 = ", 1>>2)
	println("1 &^ 2 = ", 1&^2)
	println("1 == 2 = ", 1 == 2)
	println("1 != 2 = ", 1 != 2)
	println("1 < 2 = ", 1 < 2)
	println("1 <= 2 = ", 1 <= 2)
	println("1 > 2 = ", 1 > 2)
	println("1 >= 2 = ", 1 >= 2)

	var a bool = 10 == 10
	var b bool = 20 == 10
	println("1 && 2 = ", a && b)

	println("1 || 2 = ", a || b)
	println("1 &^ 2 = ", 1&^2)
}