package main

import (
	"fmt"
	"log"

	"github.com/fmo/go-rest/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Can't load config")
	}

	fmt.Println(cfg.Env)
}
