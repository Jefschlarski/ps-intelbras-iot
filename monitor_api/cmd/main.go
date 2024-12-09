package main

import (
	"fmt"

	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules"
	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry"
	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/pkg/config"
	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/pkg/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	err := config.Load()
	if err != nil {
		panic(err)
	}

	dbconfig := config.GetDbConfig()

	// Initialize gin server engine
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // Permite qualquer origem, substitua por uma lista específica de origens, se necessário
		AllowMethods:     []string{"GET"},                                     // Métodos permitidos
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true,                                                // Permitir envio de credenciais (cookies, headers de autenticação, etc.)
	}))

	// Initialize database
	db, err := db.ConnectDB(dbconfig)
	if err != nil {
		panic(err)
	}

	// Modules
	modules := []modules.ModuleInterface{
		telemetry.NewTelemetryModule(),
	}

	v1 := server.Group("/api/v1")

	// Initialize modules
	for _, module := range modules {
		module.Init(v1, db)
	}

	// Ping route
	server.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	apiConfig := config.GetApiConfig()

	server.Run(fmt.Sprintf("%s:%d", apiConfig.Url, apiConfig.Port))
}
