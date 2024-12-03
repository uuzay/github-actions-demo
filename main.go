package main

import "fmt"

func main() {

	f1()
	f2(10)
	val := functionThatReturnsThree()
	fmt.Printf("Main replies \"it actually returned: %v\"", val)
}

func f1() {
	fmt.Printf("Function 1 speaking!\n\n")
}

func f2(val int) {
	fmt.Println("Function 2 speaking!")
	if val > 0 {
		fmt.Printf("Val is positive: %v\n\n", val)
	} else {
		fmt.Printf("Val is non-positive: %v\n\n", val)
	}
}

func functionThatReturnsThree() int {
	fmt.Println("functionThatReturnsThree speaking!")
	fmt.Printf("I shall return 3 no matter what!\n\n")

	//You only had one job...
	return 3
}
