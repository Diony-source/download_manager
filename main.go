package main

import (
	"log"
	"download_manager/handlers"
)

func main() {
	log.Println("Starting Download Manager...")
	handlers.StartCLI()
}