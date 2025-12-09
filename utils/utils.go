package utils

import (
	"log"
	"os"
)

func ReadInputFromFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading from file %s: %v\n", fileName, err)
	}
	cont := string(content)

	return cont
}
