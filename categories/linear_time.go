package main

import (
	"fmt"
	"time"
)

func linear(n int) {
	for i := 1; i <= n; i++ { // execution time increases according to variable n
		fmt.Println("loop execution:", i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	start := time.Now()
	linear(5)
	fmt.Println("time elapsed:", time.Since(start))
	start = time.Now()
	linear(10)
	fmt.Println("time elapsed:", time.Since(start))
}

