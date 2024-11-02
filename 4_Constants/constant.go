package main

import "fmt"

const marks = 30

//This is allowed
const Name string = "Hello"

//name := "Diwakar"
//Will give error := non-declaration statement outside function body

func main() {
	//marks=34 will give error because constant can not modify

	const a = 23
	fmt.Println(a)
	fmt.Println(marks)
	fmt.Println(Name)
}
