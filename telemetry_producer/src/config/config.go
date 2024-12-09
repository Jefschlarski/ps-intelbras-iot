package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *config

// config representa a estrutura de configuração do serviço
type config struct {
	Grpc  GrpcConfig
	Kafka KafkaConfig
	Log   LogConfig
}

// GrpcConfig representa a configuração do gRPC
type GrpcConfig struct {
	Port    uint
	Workers int
}

func (g GrpcConfig) String() string {
	return fmt.Sprintf("Port: %d, Workers: %d", g.Port, g.Workers)
}

// MongoConfig representa a configuração do MongoDB
type MongoConfig struct {
	Uri      string
	Port     int
	Database string
}

func (m MongoConfig) String() string {
	return fmt.Sprintf("Uri: %s, Port: %d, Database: %s", m.Uri, m.Port, m.Database)
}

// KafkaConfig representa a configuração do Kafka
type KafkaConfig struct {
	Server                    string
	Topic                     string
	QueueBufferingMaxMessages int
	QueueBufferingMaxKbytes   int
	CompressionType           string
	BatchSize                 int
}

func (k KafkaConfig) String() string {
	return fmt.Sprintf("Server: %s, Topic: %s, QueueBufferingMaxMessages: %d, QueueBufferingMaxKbytes: %d, CompressionType: %s, BatchSize: %d",
		k.Server, k.Topic, k.QueueBufferingMaxMessages, k.QueueBufferingMaxKbytes, k.CompressionType, k.BatchSize)
}

// LogConfig representa a configuração do log
type LogConfig struct {
	LogFile            string
	LogFileMaxBytes    int64
	LogFileBackupCount int
}

func (l LogConfig) String() string {
	return fmt.Sprintf("LogFile: %s, LogFileMaxBytes: %d, LogFileBackupCount: %d",
		l.LogFile, l.LogFileMaxBytes, l.LogFileBackupCount)
}

// init configura as variáveis globais com seus valores padrão
func init() {
	viper.SetDefault("grpc.port", 50051)
	viper.SetDefault("grpc.workers", 4)

	viper.SetDefault("kafka.server", "localhost:9092")
	viper.SetDefault("kafka.topic", "telemetry")
	viper.SetDefault("kafka.queue_buffering_max_messages", 10000000)
	viper.SetDefault("kafka.queue_buffering_max_kbytes", 1048576)
	viper.SetDefault("kafka.compression_type", "snappy")
	viper.SetDefault("kafka.batch_size", 10)

	viper.SetDefault("log.file", "./logs/app.log")
	viper.SetDefault("log.file_max_bytes", 50000)
	viper.SetDefault("log.file_backup_count", 5)
}

// Load carrega o arquivo de configuração
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	cfg = new(config)

	cfg.Grpc = GrpcConfig{
		Port:    viper.GetUint("grpc.port"),
		Workers: viper.GetInt("grpc.workers"),
	}

	cfg.Kafka = KafkaConfig{
		Server:                    viper.GetString("kafka.server"),
		Topic:                     viper.GetString("kafka.topic"),
		QueueBufferingMaxMessages: viper.GetInt("kafka.queue_buffering_max_messages"),
		QueueBufferingMaxKbytes:   viper.GetInt("kafka.queue_buffering_max_kbytes"),
		CompressionType:           viper.GetString("kafka.compression_type"),
		BatchSize:                 viper.GetInt("kafka.batch_size"),
	}

	cfg.Log = LogConfig{
		LogFile:            viper.GetString("log.file"),
		LogFileMaxBytes:    viper.GetInt64("log.file_max_bytes"),
		LogFileBackupCount: viper.GetInt("log.file_backup_count"),
	}

	return nil
}

// GetGrpcConfig retorna a configuração do gRPC
func GetGrpcConfig() *GrpcConfig {
	return &cfg.Grpc
}

// GetKafkaConfig retorna a configuração do Kafka
func GetKafkaConfig() *KafkaConfig {
	return &cfg.Kafka
}

// GetLogConfig retorna a configuração do log
func GetLogConfig() *LogConfig {
	return &cfg.Log
}
