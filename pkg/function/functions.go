package function

import (
	"fmt"
	"time"
)

func PrintFunctions() {

	//1: create a function
	println("1: create a function")
	//The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:

	//2: create a function which counts to 1 billion and print the time it took
	println("2: create a function which counts to 1 billion and print the time it took")

	start := time.Now()

	countToBillion()

	elapsed := time.Since(start)
	fmt.Printf("Counting to 1 billion took %s\n", elapsed)

}

func countToBillion() {
	for i := 0; i < 1000000000; i++ {
	}
}