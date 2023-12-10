package letsgo

func Functions() {

	//1: create a function
	println("1: create a function")
	func() {
		println("Hello World")
	}()

	//2: create a function with parameters
	println("2: create a function with parameters")
	func(a int, b int) {
		println(a + b)
	}(10, 20)

	//3: create a function with return value
	println("3: create a function with return value")
	z := func(a int, b int) int {
		return a + b
	}(100, 20)
	println(z)

	//4: create a function with multiple return values
	println("4: create a function with multiple return values")
	a, b := func(a int, b int) (int, int) {
		return a + b, a - b
	}(100, 200)
	println(a, b)

	//5: create a function with named return values
	println("5: create a function with named return values")
	a, b = func(a int, b int) (x int, y int) {
		x = a + b
		y = a - b
		return
	}(100, 200)
	println(a, b)

	//6: create a function with variadic parameters
	println("6: create a function with variadic parameters")
	func(a ...int) {
		for _, v := range a {
			println(v)
		}
	}(10, 20, 30, 100)

}