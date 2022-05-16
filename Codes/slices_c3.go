package main

import "fmt"

func main() {

	cards := []string{"Zero", "One", "Two", "Three", "Four"}
	//all values
	fmt.Println(cards)
	//first two
	fmt.Println(cards[:2])
	//from two
	fmt.Println(cards[2:])
}
