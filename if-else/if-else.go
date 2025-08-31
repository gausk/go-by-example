package main

import "fmt"

func main() {
	if 8%2 == 0 {
		fmt.Println("8 is even")
	} else {
		fmt.Println("8 is odd")
	}

	if num := 7; num%2 == 0 {
		fmt.Println(num, "is even")
	} else if num%3 == 0 {
		fmt.Println(num, "divisible by 3")
	} else {
		fmt.Println(num, "is odd")
	}
}
