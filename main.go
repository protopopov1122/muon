package main

import (
	"os"

	"protopopov.lv/muon"
)

func main() {
	logger := muon.NewDefaultLogger(os.Stdout)
	configuration, err := muon.LoadConfigFromEnv(logger)
	if err != nil {
		logger.Fatal.Fatalln("Failed to load service configuration due to:", err)
	}
	srv, err := muon.NewServer(logger, configuration)
	if err != nil {
		logger.Fatal.Fatalln("Failed to initialize service due to:", err)
	}
	logger.Fatal.Fatalln("Failure of service: ", srv.Start())
}
