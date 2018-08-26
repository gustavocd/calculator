package calculator

import "fmt"

// Calculate makes simple arithmetic operation
func Calculate(a, b int) string {
	return fmt.Sprintf("¡La suma de %d más %d es igual a: %d!", a, b, a+b)
}
