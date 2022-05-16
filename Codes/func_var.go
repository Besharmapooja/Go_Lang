package main

import "fmt"

func main() {
	add := func(i, j int) int {
		return i + j

	}

	sub := func(i, j int) int {
		return i - j

	}
	fmt.Println("Sum of 1+4 = ", add(1, 4))
	fmt.Println("Sub of 1-4 = ", sub(1, 4))
	fmt.Println("op of 1+4 = ", doAurth(add, 1, 4))
	fmt.Println("op of 1-4 = ", doAurth(sub, 1, 4))

	func() {
		fmt.Println("right now")
	}()

	defer func() {
		fmt.Println("defer now")
	}()
}

func add(i, j int) int {
	return i + j
}

func doAurth(op func(int, int) int, i int, j int) int {
	return (op(i, j))
}
