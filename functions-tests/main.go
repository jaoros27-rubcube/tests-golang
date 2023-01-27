package main

import "fmt"

type Car struct {
	Nome string
}

func (c Car) walk() {
	fmt.Println(c.Nome, "andou")
}

func main() {
	result1, str := sum1(10, 20)
	fmt.Println(result1, str)

	result2 := sum2(10, 20)
	fmt.Println(result2)

	result3 := sumAll(10, 20, 5, 3, 4, 6)
	fmt.Println(result3)

	result4 := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}
		return func() int {
			return res * res
		}
	}
	fmt.Println(result4(10, 20, 3, 52, 30)())

	car := Car{
		"Ferrari",
	}
	car.walk()
}

func sum1(a int, b int) (int, string) {
	return a + b, "somou!"
}

func sum2(a int, b int) (result int) {
	result = a + b
	return
}

func sumAll(x ...int) int {
	result := 0

	for _, v := range x {
		result += v
	}

	return result
}
