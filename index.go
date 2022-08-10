package main

import (
	"awesomeProject/route"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	route.CreateUrlMappings()
	route.Router.Run(":8085") // Listen and server on 0.0.0.0:8080
}
