package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(res.Body)

	respSoma, err := soma(3, 2)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(respSoma)
}

func soma(x int, y int) (int, error) {

	res := x + y

	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}
