package main

import "github.com/Cameronketchem/CEN3031-Group91/server/routes"

func main() {
	api := routes.New("server.sqlite3", "0x123456789ABCDEF", false)
	api.Start(":8080", false)
}
