package validator

import "strconv"

// IsLuhnValid verifica se o PAN passa no algoritmo de Luhn
func IsLuhnValid(pan string) bool {
	sum := 0
	alternate := false

	for i := len(pan) - 1; i >= 0; i-- {
		// Converte caractere para dígito inteiro
		n, err := strconv.Atoi(string(pan[i]))
		if err != nil {
			return false
		}

		if alternate {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alternate = !alternate
	}
	return sum%10 == 0
}
