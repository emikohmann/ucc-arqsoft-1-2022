package main

import "fmt"

type MyStruct struct {
	Name string
	Age  int
}

func main() {
	var anyString string = "Hello"
	Show(anyString)

	var anyInt int = 12345
	Show(anyInt)

	var myStruct = MyStruct{
		Name: "Emi",
		Age:  27,
	}
	Show(myStruct)
}

func Show(message interface{}) {
	fmt.Println(message)
}
