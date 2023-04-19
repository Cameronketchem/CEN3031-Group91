package main

import (
	"github.com/Cameronketchem/CEN3031-Group91/server/routes"
	"github.com/Cameronketchem/CEN3031-Group91/server/utils"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// Load environment variables.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables.
	dbpath := utils.GetVar("DB_PATH")
	secret := utils.GetVar("SECRET")
	privKey := utils.GetVar("PRIV_KEY")
	nodeAddr := utils.GetVar("NODE_ADDR")

	// Serve on localhost:8080
	api := routes.New(dbpath, secret, privKey, nodeAddr, false)
	api.Start(":8080", true)
}
