package calc

import "log"

type calculator struct{}

func BuildCalculator() calculator {
	return calculator{}
}

func (cal calculator) add(num1, num2 float64) float64 {
	return num1 + num2
}

func (cal calculator) subtract(num1, num2 float64) float64 {
	return num1 - num2
}

func (cal calculator) divide(num1, num2 float64) float64 {
	if num2 == 0 {
		log.Fatalln("Делить на 0 нельзя")
		return 0
	}
	return num1 / num2
}

func (cal calculator) multiply(num1, num2 float64) float64 {
	return num1 * num2
}

func (cal calculator) Calculate(num1, num2 float64, oper string) float64 {
	switch oper {
	case "+":
		return cal.add(num1, num2)
	case "-":
		return cal.subtract(num1, num2)
	case "*":
		return cal.multiply(num1, num2)
	case "/":
		return cal.divide(num1, num2)
	}
	return 0
}
