package main

import (
  "fmt"
  "log"
  "os"
  "github.com/joho/godotenv"
  "github.com/Cameronketchem/CEN3031-Group91/server/routes"
)

// Returns an environment variable, issues error
// if it doesnt exist.
func getVar(key string) string {
  value, exists := os.LookupEnv(key)
  if !exists {
    log.Fatal(fmt.Sprintf("Expected %s environment variable to be set", key))
  }

  return value
}

func main() {
  // Load environment variables.
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  // Retrieve environment variables.
  dbpath := getVar("DB_PATH")
  secret := getVar("SECRET")
  privKey := getVar("PRIV_KEY")
  nodeAddr := getVar("NODE_ADDR")

  // Serve on localhost:8080
  api := routes.New(dbpath, secret, privKey, nodeAddr, false)
	api.Start(":8080", true)
}
