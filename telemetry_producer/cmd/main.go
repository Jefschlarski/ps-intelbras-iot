package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/config"
	pb "github.com/Jefschlarski/ps-intelbras-iot/producer/src/grpc"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/kafka"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/utils"
	"google.golang.org/grpc"
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

	// Conecta ao Kafka
	kafkaConfig := config.GetKafkaConfig()
	logger.Info("Conectando ao Kafka")
	logger.Debug(kafkaConfig.String())
	producer, err := kafka.NewKafkaProducer(kafkaConfig)
	if err != nil {
		logger.Critical(fmt.Sprintf("Erro ao conectar ao Kafka: %v", err))
		os.Exit(1)
	}
	defer producer.Close()

	grpcConfig := config.GetGrpcConfig()
	logger.Info("Iniciando o servidor gRPC")
	logger.Debug(grpcConfig.String())

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcConfig.Port))
	if err != nil {
		logger.Critical(fmt.Sprintf("Erro ao iniciar o listener: %v", err))
		os.Exit(1)
	}

	// Cria um servidor gRPC
	grpcServer := grpc.NewServer()

	telemetry := pb.NewGrpcService(producer, grpcConfig.Workers)
	defer telemetry.GracefulStop()

	// Registra o serviço TelemetryService
	pb.RegisterTelemetryServiceServer(grpcServer, telemetry)

	// Inicia o servidor")
	if err := grpcServer.Serve(lis); err != nil {
		logger.Critical(fmt.Sprintf("Erro ao iniciar o servidor gRPC: %v", err))
		os.Exit(1)
	}
}
