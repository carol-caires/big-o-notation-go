package main

import "fmt"

func constant(n int) {
	if n > 0 {
		fmt.Println(fmt.Sprintf("input \"%d\" is positive", n))
	} else {
		fmt.Println(fmt.Sprintf("input \"%d\" is negative", n))
	}

	fmt.Println("now, we're gonna count from 1 to 10")
	for i := 1; i <= 10; i++ { // this loop does not rely on n variable, so it's O(1)
		fmt.Print(i, " ")
	}
	println("")
}

func main() {
	constant(7) // the execution time will be always the same, regardless the input size
}