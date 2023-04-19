package utils

import (
	"fmt"
	"log"
	"os"
)

// Returns an environment variable, issues error if it doesnt exist.
func GetVar(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal(fmt.Sprintf("Expected %s environment variable to be set", key))
	}

	return value
}
