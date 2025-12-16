package emv

import "testing"

func TestProcess(t *testing.T) {
	t.Run("Cartão Expirado", func(t *testing.T) {
		input := Transaction{PAN: "4539123456789012", Expiry: "100101"}
		res, err := Process(input)
		if res != "REJEITADA" || err.Error() != "cartão expirado" {
			t.Errorf("Esperava erro de expiração, obteve: %v", err)
		}
	})
}
