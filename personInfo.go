package main

import "fmt"


type Adress struct {
	City string
	Street string
}


type Person struct {
	Name string
	Lastname string
	Age int64
	Adress Adress
}


func (p Person) Description() string {
	return fmt.Sprintf("%s %s, %d godina, živi u: %s, %s.", p.Name, p.Lastname, p.Age, p.Adress.City, p.Adress.Street)
}


func main() {

	adress := Adress{
		City: "Mostar",
		Street: "Ulica kralja Tomislava 1",
	}

	person := Person{
		Name: "Pero",
		Lastname: "Perić",
		Age: 20,
		Adress: adress,
	}

	fmt.Println(person.Description())

}
