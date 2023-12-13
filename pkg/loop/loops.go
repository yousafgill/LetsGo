package loop

func PrintLoops() {
	//1:  for loop
	println("for loop")
	for i := 0; i < 10; i++ {
		println(i)
	}

	//2: for loop with condition
	println("for loop with condition")
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	//3: for loop with break
	println("for loop with break")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		println(i)
	}

	//4: for loop with continue
	println("for loop with continue")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		println(i)
	}

	//5: for loop with range
	println("for loop with range")
	s := []int{1, 2, 3}
	for i, v := range s {
		println(i, v)
	}

	//6: for loop with range
	println("for loop with range")
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		println(k, v)
	}

	//7: for loop with range
	println("for loop with range")
	s2 := "hello"
	for i, v := range s2 {
		println(i, v)
	}

}