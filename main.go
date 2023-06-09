package main

import (
	"fmt"

	"calculator/calc"
)

func main (){
	fmt.Println("Add:", calc.Add(5, 5))
	fmt.Println("Sub:", calc.Sub(5, 5))
	fmt.Println("Multiple:", calc.Mult(5, 5))
	fmt.Println("Divide:", calc.Div(5, 5))
}