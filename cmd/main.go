package main

import (
	"log"

	"github.com/MdSadiqMd/Quantum-Cart-Backend/internal/api"
	"github.com/MdSadiqMd/Quantum-Cart-Backend/packages/config"
)

func main() {
	cfg, err := config.SetupEnv()
	if err != nil {
		log.Fatalf("error in config: %v", err)
	}

	api.StartServer(cfg)
}
