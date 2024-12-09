package kafka

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/config"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/dto"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/service"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaConsumer representa um consumidor de mensagens Kafka
type kafkaConsumer struct {
	consumer *kafka.Consumer
	topic    string
	wg       sync.WaitGroup
	service  service.TelemetryServiceInterface
	log      utils.Logger
}

// NewKafkaConsumer cria um novo consumidor Kafka
func NewKafkaConsumer(config *config.KafkaConfig, service service.TelemetryServiceInterface) (*kafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  config.Server,
		"group.id":           config.GroupID,
		"auto.offset.reset":  "latest",
		"enable.auto.commit": "false",
	})
	if err != nil {
		return nil, err
	}

	err = consumer.SubscribeTopics([]string{config.Topic}, nil)
	if err != nil {
		return nil, err
	}

	return &kafkaConsumer{
		consumer: consumer,
		topic:    config.Topic,
		service:  service,
		log:      *utils.GetLoggerInstance(),
	}, nil
}

// Consume consome mensagens do Kafka em um loop
func (c *kafkaConsumer) Consume() error {
	c.wg.Add(1)
	defer c.wg.Done()

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			c.log.Error(fmt.Sprintf("Erro ao ler mensagem: %v", err))
			continue
		}

		messageValue := msg.Value

		c.log.Info(fmt.Sprintf("Mensagem recebida: %s", string(messageValue)))

		var telemetryEvent dto.TelemetryEventDto
		err = json.Unmarshal(messageValue, &telemetryEvent)
		if err != nil {
			c.log.Error(fmt.Sprintf("Erro ao deserializar mensagem: %v", err))
			continue
		}

		if err := telemetryEvent.Validate(); err != nil {
			c.log.Error(fmt.Sprintf("Erro de validação: %v", err))
			continue
		}

		err = c.service.Save(&telemetryEvent)
		if err != nil {
			c.log.Error(fmt.Sprintf("Erro ao salvar dados: %v", err))
			continue
		}

		_, err = c.consumer.CommitMessage(msg)
		if err != nil {
			c.log.Error(fmt.Sprintf("Erro ao commitar mensagem: %v", err))
			continue
		}

	}
}

// Close fecha o consumidor Kafka
func (c *kafkaConsumer) Close() {
	c.consumer.Close()
	c.wg.Wait()
}
