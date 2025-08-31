package main

import "fmt"

func main() {
	i := 1

	for i <= 3 {
		fmt.Println("Simple loop", i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println("C like loop", j)
	}

	for i := range 5 {
		fmt.Println("Range loop", i)
	}

	for {
		fmt.Println("Break loop")
		break
	}

	for n := range 10 {
		if n%2 == 0 {
			continue
		}
		fmt.Println("Continue loop", n)
	}
}
