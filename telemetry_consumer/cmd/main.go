package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/config"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/db"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/kafka"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/repo"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/service"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/utils"
)

func main() {

	// Carrega o arquivo de configuração
	err := config.Load()
	if err != nil {
		log.Println(err)
		log.Fatalf("Falha ao carregar o arquivo de configuração")
	}

	// Inicializa o logger
	logger := utils.NewLogger(config.GetLogConfig())

	logger.Info("Iniciando Serviço")

	// Conecta ao banco
	dbConfig := config.GetDBConfig()
	logger.Info("Conectando ao banco de dados")
	logger.Debug(dbConfig.String())
	db, err := db.ConnectDB(dbConfig)
	if err != nil {
		logger.Critical(fmt.Sprintf("Erro ao conectar ao banco de dados: %v", err))
		os.Exit(1)
	}
	defer db.Close()

	// Iniciar o repositorio
	repo := repo.NewTelemetryRepo(db)

	// Iniciar o service
	service := service.NewTelemetryService(repo)

	// Conecta ao Kafka
	kafkaConfig := config.GetKafkaConfig()
	logger.Info("Conectando ao Kafka")
	logger.Debug(kafkaConfig.String())
	consumer, err := kafka.NewKafkaConsumer(kafkaConfig, service)
	if err != nil {
		logger.Critical(fmt.Sprintf("Erro ao conectar ao Kafka: %v", err))
		os.Exit(1)
	}
	consumer.Consume()
	defer consumer.Close()
}
