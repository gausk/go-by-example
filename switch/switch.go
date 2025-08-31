package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2

	switch i {
	case 1:
		fmt.Println(i, "should be written as one")
	case 2:
		fmt.Println(i, "should be written as two")
	case 3:
		fmt.Println(i, "should be written as three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hello Weekends")
	default:
		fmt.Println("Working day Bro")
	}

	t := time.Now().Hour()

	switch {
	case t < 12:
		fmt.Println("Good Morning")
	default:
		fmt.Println("Good Afternoon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a int")
		default:
			fmt.Printf("Don't know a type: %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(2)
	whatAmI("alpha")
}
