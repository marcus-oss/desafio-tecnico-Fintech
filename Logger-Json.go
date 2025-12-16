package storage

import (
	"encoding/json"
	"os"
)

type LogEntry struct {
	Timestamp string `json:"timestamp"`
	Status    string `json:"status"`
	Message   string `json:"message"`
}

// SaveLog grava o resultado em um arquivo .json
func SaveLog(entry LogEntry) error {
	file, _ := os.OpenFile("transactions.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(entry)
}
