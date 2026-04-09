package main

import "fmt"

func main() {
	fmt.Println("==== CALCULARORA CIENTÍFICA v1.0 ====")
	var num1, num2 float64
	fmt.Println("Ingresa el primer número:")
	fmt.Scanln(&num1)
	fmt.Println("Ingresa el segundo número:")
	fmt.Scanln(&num2)
	fmt.Println("Ingresa la operación (+, -, *, /):")
	var operacion string
	fmt.Scanln(&operacion)
	switch operacion {
	case "+":
		fmt.Printf("Resultado: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Resultado: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Resultado: %.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("Resultado: %.2f\n", num1/num2)
		} else {
			fmt.Println("Error: División por cero")
		}
	default:
		fmt.Println("Operación no válida")
	}
}
