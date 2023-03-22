package calc

import "log"

type calculator struct{}

func BuildCalculator() calculator {
	return calculator{}
}

func (calc calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

func (calc calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func (calc calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		log.Fatalln("На 0 делить нельзя")
		return 0
	}
	return num1 / num2
}

func (calc calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func (calc calculator) Calculate(num1, num2 float64, oper string) float64 {
	switch oper {
	case "+":
		return calc.add(num1, num2)
	case "-":
		return calc.subtract(num1, num2)
	case "*":
		return calc.multiply(num1, num2)
	case "/":
		return calc.divide(num1, num2)
	}
	return 0
}
