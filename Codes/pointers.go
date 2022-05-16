package main

import "fmt"

type person struct {
	name string
	zip  int
}

func main() {
	p := person{"Pooja", 1234}
	fmt.Println("Date", p)
	personPointer := &p
	fmt.Println("Address", *personPointer)
	p.updatePersonReceiver("Pooja")
	fmt.Println("Updated", p)
	personPointer.updatePersonReceiver("suuu")
	fmt.Println("updatePersonReceiver", p)
}

func (p person) updatePersonReceiver(name string) {
	p.name = name

}
func (p *person) updatePersonRecevier(name string) {
	p.name = name

}

// reference type-> slice map
// value type-> int sting array
