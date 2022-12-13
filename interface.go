package main

import "fmt"

type (
	HasName interface {
		GetName() string
	}

	Person struct {
		Name string
	}

	Animal struct {
		Name string
	}
)

func SayHello(hasName HasName) {
	fmt.Println("Hello my name is", hasName.GetName())
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var name Person
	name.Name = "Budi"
	SayHello(name)

	name_animal := Animal{
		Name: "Kucing",
	}
	SayHello(name_animal)
}
