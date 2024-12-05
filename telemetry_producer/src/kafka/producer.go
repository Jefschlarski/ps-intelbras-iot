package kafka

import (
	"github.com/IBM/sarama"
)

// KafkaProducer representa um produtor de mensagens Kafka
type KafkaProducer struct {
	produtor sarama.SyncProducer
	topic    string
}

// NewKafkaProducer cria um novo produtor Kafka
func NewKafkaProducer(brokers []string, topic string) (*KafkaProducer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_2_1_0
	config.Producer.Return.Successes = true

	produtor, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &KafkaProducer{
		produtor: produtor,
		topic:    topic,
	}, nil
}

// Produce envia uma mensagem para o tópico
func (k *KafkaProducer) Produce(mensagem string) error {
	msg := &sarama.ProducerMessage{
		Topic: k.topic,
		Value: sarama.StringEncoder(mensagem),
	}

	_, _, err := k.produtor.SendMessage(msg)
	return err
}

// Close fecha a conexão com o Kafka
func (k *KafkaProducer) Close() error {
	return k.produtor.Close()
}
