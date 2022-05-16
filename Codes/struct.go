package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
	addr      address
}

type address struct {
	city    string
	zipcode int
}

func main() {
	//	var p person = person{"Pooja", "Sharma", 12}
	//p := person{"Pooja", "Sharma", 12}
	p := person{firstName: "Pooja", lastName: "Sharma"}
	//p:= person{}
	// updating the value
	p.age = 12
	//fmt.Println(p)
	fmt.Printf("%+v", p)
	np := person{
		lastName:  "Sharma",
		firstName: "Pooja",
		age:       12,
		addr: address{
			city: "Terre",
		},
	}
	np.addr.zipcode = 47807
	np.updateAddress("dallas", 47444)

	fmt.Printf("%+v", p)

}
func (p person) updateAddress(city string, zip int) {
	p.addr.city = city
	p.addr.zipcode = zip
	fmt.Printf("Inside address:: %+v", p)
}
