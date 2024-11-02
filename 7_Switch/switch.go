package main

import (
	"fmt"
)

func main() {
	//Simple Switch
	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("It's a ", i)
	// case 2:
	// 	fmt.Println("It's a", i)
	// case 3:
	// 	fmt.Println("It's a ", i)
	// default:
	// 	fmt.Println("Other", i)
	// }

	//Multiple Condition Switch
	//time.Now().Weekday() will return today's day
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's weekend")
	// default:
	// 	fmt.Println("It's workday", time.Now().Weekday())
	// }

	//type switch
	// forNoReason := func(i interface{}) {
	// 	switch i.(type) {
	// 	case string:
	// 		fmt.Println("It's a String")
	// 	case bool:
	// 		fmt.Println("It's a Boolean")
	// 	case int:
	// 		fmt.Println("It's a Number")
	// 	default:
	// 		fmt.Println("Other")
	// 	}
	// }
	// forNoReason(33.6)
}
