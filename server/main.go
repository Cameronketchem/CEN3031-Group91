package main

import (
	"github.com/Cameronketchem/CEN3031-Group91/server/routes"
)

func main() {
	api := routes.New("server.sqlite3", false)
	api.Start(":8080", true)
}
