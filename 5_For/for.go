package main

import "fmt"

func main() {
	//There is no while loop in goLang only for loop exists
	//You can also use break & continue in this goLang Language and it works just like in any other language

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }

	//Here is how you can make write for loop as while loop
	// i:=0
	// for i<5{
	// 	fmt.Println(i)
	// 	i=i+1
	// }

	//Infinite loop
	// i:=1
	// for {
	// 	fmt.Println(i)
	// }

	//Using Range in for loop
	for i := range 8 {
		fmt.Println(i)
	}

}
