package main

import (
	"time"
)

type EncryptedResponse struct {
	EncryptedAESKey string `json:"encryptedAESKey"`
	EncryptedData   string `json:"encryptedData"`
	Signature       string `json:"signature"`
}

func main() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			run()
		}
	}
}
