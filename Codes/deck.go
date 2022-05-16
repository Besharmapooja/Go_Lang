package main

import "fmt"

type deck []string

//function with receiver arg aka Method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

//func with return arg
func createNewDeck() deck {
	suits := []string{"Spades", "Diamod", "Clubs", "Hearts"}
	values := []string{"Ace", "2", "3", "4"} // make it 13

	newDeck := deck{}
	//loop using range
	for _, suit := range suits {

		for j := 0; j < len(values); j++ {
			newDeck = append(newDeck, values[j]+" of "+suit)
		}
	}

	return newDeck
}
