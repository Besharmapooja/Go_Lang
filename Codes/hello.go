package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	msg := "Hi there"
	fmt.Println(msg)
	msg = "10"
	myBool := true
	myInt := 20
	var myFloat float32 = 23.65

	fmt.Printf("type is %T, value is %v \n", msg, msg)
	fmt.Printf("type is %T, value is %v \n", myBool, myBool)
	fmt.Printf("type is %T, value is %v \n", myInt, myInt)
	fmt.Printf("type is %T, value is %v \n", myFloat, myFloat)
}
