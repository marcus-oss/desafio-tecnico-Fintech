package emv

import (
	"emv-transaction/internal/validator"
	"errors"
	"math/rand"
	"time"
)

// Transaction representa os dados de uma tentativa de venda
type Transaction struct {
	PAN    string // Tag 5A
	Expiry string // Tag 5F24 (YYMMDD)
	CVM    string // Tag 9F34
}

// Process analisa os dados e decide se envia para o gateway
func Process(data Transaction) (string, error) {
	// 1. Validação de Tamanho do PAN
	if len(data.PAN) < 13 || len(data.PAN) > 19 {
		return "REJEITADA", errors.New("tamanho de PAN inválido")
	}

	// 2. Validação Luhn
	if !validator.IsLuhnValid(data.PAN) {
		return "REJEITADA", errors.New("falha no algoritmo de Luhn")
	}

	// 3. Validação de Data (Simplificada YYMMDD)
	currentDate := time.Now().Format("060102") // Formato YYMMDD
	if data.Expiry < currentDate {
		return "REJEITADA", errors.New("cartão expirado")
	}

	// 4. Simulação de Gateway (Mock)
	return callGateway(), nil
}

func callGateway() string {
	// Simula 80% de chance de aprovação
	if rand.Float32() > 0.2 {
		return "APROVADA"
	}
	return "REJEITADA: Saldo Insuficiente"
}
