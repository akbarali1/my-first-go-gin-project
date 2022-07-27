package main

import (
	"awesomeProject/mappings"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mappings.CreateUrlMappings()
	mappings.Router.Run(":8085") // Listen and server on 0.0.0.0:8080
}
