package main

import (
	"fmt"
	"variables/math"
)

func main() {
	var a float64 = 2.54
	var b float64 = 3.96
	resultado := math.Add(a, b)
	fmt.Println(resultado)
}
