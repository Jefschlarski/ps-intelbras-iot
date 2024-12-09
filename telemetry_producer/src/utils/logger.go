package utils

import (
	"log"

	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	logger *log.Logger
}

var instance *Logger

// Init initializes the logger with a rotating file handler and a console handler.
func NewLogger(config *config.LogConfig) *Logger {

	if instance == nil {

		// Set up the rotating file handler
		lumberjackLogger := &lumberjack.Logger{
			Filename:   config.LogFile,
			MaxSize:    int(config.LogFileMaxBytes / (1024 * 1024)),
			MaxBackups: config.LogFileBackupCount,
			MaxAge:     28,
			Compress:   true,
		}

		// Create a new logger
		instance = &Logger{
			logger: log.New(lumberjackLogger, "", log.LstdFlags),
		}

	}

	return instance

}

func GetLoggerInstance() *Logger {
	return instance
}

func (l *Logger) Debug(message string) {
	log.Println(message)
	l.logger.Println("🐞 " + message)
}

func (l *Logger) Info(message string) {
	log.Println(message)
	l.logger.Println(message)
}

func (l *Logger) Warning(message string) {
	log.Println(message)
	l.logger.Println("⚠️ " + message)
}

func (l *Logger) Error(message string) {
	log.Println(message)
	l.logger.Println("❌ " + message)
}

func (l *Logger) Critical(message string) {
	log.Println(message)
	l.logger.Println("🆘 " + message)
}
