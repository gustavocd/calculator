package calculator

import "fmt"

// Calculate makes simple arithmetic operations
// such as add, substract, multiply and divide
func Calculate(operation string, a, b int) string {
	switch operation {
	case "add":
		return fmt.Sprintf("¡La suma de %d más %d es igual a: %d!", a, b, a+b)
	case "substract":
		return fmt.Sprintf("¡La resta de %d menos %d es igual a: %d!", a, b, a-b)
	case "multiply":
		return fmt.Sprintf("¡La multiplicación de %d por %d es igual a: %d!", a, b, a*b)
	case "divide":
		return fmt.Sprintf("¡La división de %d entre %d es igual a: %d!", a, b, a/b)
	default:
		return fmt.Sprintf("¡No se conoce esa operación 😕!")
	}
}
