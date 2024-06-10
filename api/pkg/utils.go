package pkg

import (
	"crypto/rand"
	"encoding/hex"
	"log"
)

func Int32Ptr(i int32) *int32 { return &i }

func GenerateUID() string {
	randomBytes := make([]byte, 16) // 16 bytes = 128 bits
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random UID: %v", err)
	}
	return hex.EncodeToString(randomBytes)
}
