package main

import "fmt"

type Person struct {
	Name    string
	Age     int
	Address string
}

func (p Person) PrintInfo() {
	fmt.Printf("Name: %s\n Age: %d\n Address: %s\n", p.Name, p.Age, p.Address)
}
func main() {
	var orang1 = Person{
		Name:    "Mingyu",
		Age:     25,
		Address: "123 Gangnam",
	}
	var orang2 = Person{
		Name:    "Joshua",
		Age:     26,
		Address: "456 Los Angeles",
	}
	orang1.PrintInfo()
	orang2.PrintInfo()
}
