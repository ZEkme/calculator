package main

import (
	"fmt"
	"log"
	"module5/calc"
)

func isOper(oper string) bool {
	switch oper {
	case "+", "-", "*", "/":
		return true
	}
	return false
}

func main() {
	fmt.Print("Введите операцию: ")
	var oper string // operator
	fmt.Scanln(&oper)
	if !isOper(oper) {
		log.Fatalln("Недопустимая операция")
		return
	}

	var num1, num2 float64
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print("Введите второе число: ")
	_, err2 := fmt.Scanln(&num2)
	if err2 != nil {
		log.Fatalln(err)
	}

	cal := calc.BuildCalculator()
	fmt.Println(cal.Calculate(num1, num2, oper))

}
