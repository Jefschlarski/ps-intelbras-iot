package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *config

// config representa a estrutura de configuração do serviço
type config struct {
	Db    DBConfig
	Kafka KafkaConfig
	Log   LogConfig
}

// DBConfig representa a configuração do banco de dados
type DBConfig struct {
	Host     string
	Port     int
	User     string
	Pass     string
	Database string
	Drive    string
}

func (d DBConfig) String() string {
	return fmt.Sprintf("Host: %s, Port: %d, User: %s, Pass: %s, Database: %s, Drive: %s",
		d.Host, d.Port, d.User, d.Pass, d.Database, d.Drive)
}

// KafkaConfig representa a configuração do Kafka
type KafkaConfig struct {
	Server    string
	Topic     string
	GroupID   string
	BatchSize int
}

func (k KafkaConfig) String() string {
	return fmt.Sprintf("Server: %s, Topic: %s, GroupID: %s, BatchSize: %d",
		k.Server, k.Topic, k.GroupID, k.BatchSize)
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

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5150)
	viper.SetDefault("database.user", "postgres")
	viper.SetDefault("database.pass", "postgres")
	viper.SetDefault("database.name", "postgres")
	viper.SetDefault("database.drive", "postgres")

	viper.SetDefault("kafka.server", "localhost:9092")
	viper.SetDefault("kafka.topic", "telemetry")
	viper.SetDefault("kafka.group_id", "telemetry-1")
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

	cfg.Db = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
		Drive:    viper.GetString("database.drive"),
	}
	cfg.Kafka = KafkaConfig{
		Server:    viper.GetString("kafka.server"),
		Topic:     viper.GetString("kafka.topic"),
		GroupID:   viper.GetString("kafka.group_id"),
		BatchSize: viper.GetInt("kafka.batch_size"),
	}

	cfg.Log = LogConfig{
		LogFile:            viper.GetString("log.file"),
		LogFileMaxBytes:    viper.GetInt64("log.file_max_bytes"),
		LogFileBackupCount: viper.GetInt("log.file_backup_count"),
	}

	return nil
}

// GetDBConfig retorna a configuração do banco de dados
func GetDBConfig() *DBConfig {
	return &cfg.Db
}

// GetKafkaConfig retorna a configuração do Kafka
func GetKafkaConfig() *KafkaConfig {
	return &cfg.Kafka
}

// GetLogConfig retorna a configuração do log
func GetLogConfig() *LogConfig {
	return &cfg.Log
}
