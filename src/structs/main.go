package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// Different way to instantiate the struct
	// luca := person{"Luca", "Pau"}
	// var luca person
	// luca.firstName = "Luca"
	// luca.lastName = "Pau"
	luca := person{
		firstName: "Luca",
		lastName:  "Pau",
		contactInfo: contactInfo{
			email:   "luca@gmail.com",
			zipCode: 40132,
		},
	}

	// lucaPointer := &luca
	// lucaPointer.updateName("LucaChangedWithPointer")
	// shortcaut to pointer, when the receiver has a pointer, go will automatically pass a pointer
	// without explicitly create a variable pointer and call the method
	luca.updateName("LucaChangedWithPointer")
	luca.print()
}

// receiver that accept pointers in order to change the value
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
