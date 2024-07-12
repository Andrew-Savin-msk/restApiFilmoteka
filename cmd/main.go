package main

import (
	"log"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/apiserver"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/config"
)

func main() {
	cfg := config.Load()
	err := apiserver.Start(cfg)
	if err != nil {
		log.Fatalf("unable to start server. ended with error: %v", err)
	}
}
