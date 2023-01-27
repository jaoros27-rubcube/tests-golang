package main

import "fmt"

type Car struct {
	Name string
}

func (c Car) walk1() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func (c *Car) walk2() {
	c.Name = "BMW"
	fmt.Println(c.Name, "andou")
}

func main() {
	a := 10
	fmt.Println(&a)

	var pointer *int = &a
	*pointer = 50
	fmt.Println(a)

	variable := 10
	abc(&variable)

	fmt.Println(variable)

	car1 := Car{
		Name: "Ka",
	}
	car1.walk1()
	fmt.Println(car1.Name)

	car1.walk2()
	fmt.Println(car1.Name)
}

func abc(a *int) {
	*a = 200
}
