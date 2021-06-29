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

		//validando por tipo
		ce, ok := err.(*custom_error.CustomError)
		if ok {
			switch ce.Type.Codigo {
			case custom_error.Bussines_Error_1:
				fmt.Println("Error 1")
			case custom_error.Bussines_Error_2:
				fmt.Println("Error 2")
			default:
				fmt.Println("Não encontrado.")
			}
		}
	}
	fmt.Println(s)
}