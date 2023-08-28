package main

import (
	mysql "gc-hexa-go/internal/database/mysql"
	server "gc-hexa-go/pkg/http"
)

func main() {
	mysql.Connect()

	server.Serve()
}
