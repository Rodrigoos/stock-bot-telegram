package utils

import (
	"fmt"
	"strings"
)

// FormatBRL formata um valor float64 como "R$ 1.234,56"
func FormatBRL(value float64) string {
	s := fmt.Sprintf("%.2f", value)
	parts := strings.Split(s, ".")
	inteiro := parts[0]
	decimal := parts[1]

	var withDots []byte
	for i, j := len(inteiro)-1, 0; i >= 0; i, j = i-1, j+1 {
		withDots = append(withDots, inteiro[i])
		if j%3 == 2 && i != 0 {
			withDots = append(withDots, '.')
		}
	}

	for i, j := 0, len(withDots)-1; i < j; i, j = i+1, j-1 {
		withDots[i], withDots[j] = withDots[j], withDots[i]
	}

	return "R$ " + string(withDots) + "," + decimal
}
