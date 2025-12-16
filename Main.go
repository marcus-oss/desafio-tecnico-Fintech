package main

import (
	"emv-transaction/internal/emv"
	"emv-transaction/internal/storage"
	"fmt"
	"time"
)

func main() {
	// Dados simulados vindos do terminal (Chip)
	input := emv.Transaction{
		PAN:    "4539123456789012",
		Expiry: "281231",
		CVM:    "020103",
	}

	fmt.Println("Processando transação...")
	
	result, err := emv.Process(input)
	
	msg := "Sucesso"
	if err != nil {
		msg = err.Error()
	}

	// Salva log
	storage.SaveLog(storage.LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Status:    result,
		Message:   msg,
	})

	fmt.Printf("Resultado: %s | Motivo: %s\n", result, msg)
}
