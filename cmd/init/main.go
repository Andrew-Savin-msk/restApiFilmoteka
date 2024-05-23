package main

import (
	"fmt"
	"log"

	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/app/apiserver"
	"github.com/Andrew-Savin-msk/rest-api-filmoteka/internal/app/config"
)

func main() {
	var cfg config.Config
	cfg.Load()
	fmt.Println(cfg)
	err := apiserver.Start(&cfg)
	if err != nil {
		log.Fatal("Unable to stert server. Ended with error: ", err)
	}
}
