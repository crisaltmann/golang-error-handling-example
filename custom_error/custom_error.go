package custom_error

import "fmt"

type Type string

const (
	Bussines_Error_1 Type = "SAMPLE-001"
	Bussines_Error_2 Type = "SAMPLE-002"
)
type CustomError struct {
	Type 	 ErrorType
	Source   error
}

type ErrorType struct {
	Codigo Type
}

func (c *CustomError) Error() string {
	return fmt.Sprintf("[%s] %s", c.Type, c.Source.Error())
}
