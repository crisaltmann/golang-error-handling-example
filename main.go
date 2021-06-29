package main

import (
	"errors"
	"fmt"
	"github.com/crisaltmann/golang-error-handling-example/custom_error"
)



func exampleOne() (string, error) {
	err := errors.New("Exemplo de erro")
	return "", &custom_error.CustomError{
		Type:   custom_error.ErrorType{Codigo: custom_error.Bussines_Error_1},
		Source: err,
	}
}

func main() {
	s, err := exampleOne()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(s)
}