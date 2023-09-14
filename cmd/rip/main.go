package main

import (
	"RIP/internal/api"
	"log"
)

func main() {
	log.Println("Application starts!")
	api.StartServer()
	log.Println("Application terminated!")
}
