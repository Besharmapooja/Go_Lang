package main

import "fmt"

type Deck []string
type person struct {
	name string
	age  int
}

func main() {
	p := person{name: "Poo", age: 25}
	cards := Deck{"Ace", "Two"}
	//	fmt.Println(p, cards)
	cards.Print()
	p.print()
	fmt.Println(p.changeName("Pooja"))
}

func (d Deck) Print() {
	fmt.Println(d)

}

func (p person) print() {
	fmt.Println(p)
}

// need to compl
func (p person) changeName(newName string) person {
	p.name = newName
	return p
}
