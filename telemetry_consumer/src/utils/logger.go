package utils

import (
	"log"

	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *log.Logger
}

var instance *Logger

// NewLogger cria um novo logger
func NewLogger(config *config.LogConfig) *Logger {

	if instance == nil {

		lumberjackLogger := &lumberjack.Logger{
			Filename:   config.LogFile,
			MaxSize:    int(config.LogFileMaxBytes / (1024 * 1024)),
			MaxBackups: config.LogFileBackupCount,
			MaxAge:     28,
			Compress:   true,
		}

		instance = &Logger{
			logger: log.New(lumberjackLogger, "", log.LstdFlags),
		}

	}

	return instance

}

// GetLoggerInstance retorna a instância do logger
func GetLoggerInstance() *Logger {
	return instance
}

// Debug imprime uma mensagem de debug
func (l *Logger) Debug(message string) {
	log.Println(message)
	l.logger.Println("🐞 " + message)
}

// Info imprime uma mensagem de informação
func (l *Logger) Info(message string) {
	log.Println(message)
	l.logger.Println(message)
}

// Warning imprime uma mensagem de alerta
func (l *Logger) Warning(message string) {
	log.Println(message)
	l.logger.Println("⚠️ " + message)
}

// Error imprime uma mensagem de erro
func (l *Logger) Error(message string) {
	log.Println(message)
	l.logger.Println("❌ " + message)
}

// Critical imprime uma mensagem critica
func (l *Logger) Critical(message string) {
	log.Println(message)
	l.logger.Println("🆘 " + message)
}
