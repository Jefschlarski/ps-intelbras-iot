package kafka

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/config"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/utils"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaProducer representa um produtor de mensagens Kafka com suporte a batch
type kafkaProducer struct {
	producer  *kafka.Producer
	topic     string
	batchSize int
	buffer    []*kafka.Message
	mutex     sync.Mutex
	flushChan chan struct{}
	log       utils.Logger
}

// NewKafkaProducer cria um novo produtor Kafka com suporte a batch
func NewKafkaProducer(config *config.KafkaConfig) (*kafkaProducer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers":            config.Server,
		"queue.buffering.max.messages": config.QueueBufferingMaxMessages, // Número máximo de mensagens no buffer
		"queue.buffering.max.kbytes":   config.QueueBufferingMaxKbytes,   // Tamanho máximo do buffer em KB
		"compression.type":             config.CompressionType,           // Compressão para otimizar o envio
	})
	if err != nil {
		return nil, err
	}

	k := &kafkaProducer{
		producer:  producer,
		topic:     config.Topic,
		batchSize: config.BatchSize,
		buffer:    make([]*kafka.Message, 0, config.BatchSize),
		flushChan: make(chan struct{}),
		log:       *utils.GetLoggerInstance(),
	}

	// Inicia o loop de flush por tempo
	go k.startBatchFlusher(2 * time.Second)

	return k, nil
}

// Produce adiciona uma mensagem ao buffer e faz flush quando o batch estiver cheio
func (k *kafkaProducer) Produce(message string, partition int32) error {
	k.mutex.Lock()
	defer k.mutex.Unlock()

	// Adiciona a mensagem ao buffer
	k.buffer = append(k.buffer, &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &k.topic, Partition: partition},
		Value:          []byte(message),
	})

	// Verifica se o batch está cheio
	if len(k.buffer) >= k.batchSize {
		return k.flush()
	}

	return nil
}

// startBatchFlusher inicia o flush automático a cada intervalo
func (k *kafkaProducer) startBatchFlusher(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Flush automático por tempo
			k.mutex.Lock()
			if len(k.buffer) > 0 {
				if err := k.flush(); err != nil {
					log.Printf("Erro ao fazer flush: %v", err)
				}
			}
			k.mutex.Unlock()
		case <-k.flushChan:
			// Sai do loop quando o producer é fechado
			return
		}
	}
}

// flush envia todas as mensagens acumuladas no buffer
func (k *kafkaProducer) flush() error {
	if len(k.buffer) == 0 {
		return nil
	}

	var failedMessages []*kafka.Message

	// Envia as mensagens em batch
	for _, msg := range k.buffer {
		err := k.producer.Produce(msg, nil)
		if err != nil {
			k.log.Error(fmt.Sprintf("Erro ao enviar mensagem: %v", err))
			failedMessages = append(failedMessages, msg) // Armazena mensagens que falharam
		}
	}

	k.log.Info(fmt.Sprintf("Mensagens enviadas: %d", len(k.buffer)-len(failedMessages)))

	// Atualiza o buffer apenas com mensagens que falharam
	k.buffer = failedMessages

	return nil
}

// Close fecha o producer Kafka e garante que o buffer restante seja enviado
func (k *kafkaProducer) Close() {
	// Finaliza o flush automático
	close(k.flushChan)

	// Envia as mensagens restantes
	k.mutex.Lock()
	if len(k.buffer) > 0 {
		if err := k.flush(); err != nil {
			log.Printf("Erro ao fazer flush final: %v", err)
		}
	}
	k.mutex.Unlock()

	// Espera o envio de todas as mensagens pendentes
	k.producer.Flush(5000) // Aguarda até 5 segundos

	// Fecha o producer
	k.producer.Close()
}
