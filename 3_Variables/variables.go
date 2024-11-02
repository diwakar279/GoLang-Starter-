package main

import "fmt"

func main() {

	// Variables can be created by 2 ways-
	// 1. Using var keyword

	//var name Integer = 23
	//It will give error because in golang declared and not used variable will give you error
	var Brother string ="pal"
	fmt.Println(Brother)

	var isAdult bool = true
	fmt.Println(isAdult)

	//var name
	//This type of declaration will also give error
	
	var Marks = 34.5
	//or
	//var Marks float= 34.5 is same
	fmt.Println(Marks)

	// 2. Short Variable Declaration

	Price:= 234
	fmt.Println(Price)

	//Please don't confuse in between := and = as := is a declaration and = is assignment
	
}
