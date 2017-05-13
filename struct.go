package main

import "fmt"

type Person struct {
	Name    string
	Gender  string
	Address string
	Age     uint8
}

func (person *Person) Grow() {
	person.Age++
}

func (person *Person) Move(newAddress string) string {
	oldAddress := person.Address
	person.Address = newAddress
	return oldAddress
}

func main() {
	p := Person{"linjie", "male", "shanghai", 25}
	oldAddress := p.Move("beijing")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
}
