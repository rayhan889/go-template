package main

import (
	"fmt"

	"github.com/rayhan889/go-template/internal/config"
)

func main() {
    viperConfig := config.NewViper()
    log := config.NewLogger(viperConfig)
    db := config.NewDatabase(viperConfig, log)
    validate := config.NewValidator(viperConfig)
    app := config.NewFiber(viperConfig)

    config.Bootstrap(&config.BootstrapConfig{
        DB          : db,
        App         : app,
        Validator   : validate,
        Config      : viperConfig,
        Log         : log,
    })

    webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
