package main

import "fmt"

func main() {
	for i := 0; i <= 1000; i += 1 {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			// odd
		}
		fmt.Println(i)
	}
	fmt.Println("Last")
}
