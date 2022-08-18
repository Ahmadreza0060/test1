package main

import "fmt"

type person struct {
	name   string
	family string
	age    int
}

func main() {
	person1 := func(firstname string, lastname string) person {
		return person{
			name:   "ali",
			family: "rahimi",
			age:    20,
		}

	}
	person2 := person1("ali", "hasan")
	person2.Getage()
	fmt.Println(person2.age)
	fmt.Println(person2.Getage())
}
func (P *person) Getage() int {
	return P.age
}
func Getnewperson() person {
	return person{
		name:   "arash",
		family: "rahimi",
		age:    20,
	}
}
