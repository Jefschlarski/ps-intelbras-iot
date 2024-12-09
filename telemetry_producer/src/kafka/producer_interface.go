package kafka

// ProducerInterface representa a interface que deve ser implementada por um produtor de mensagens
type ProducerInterface interface {
	Produce(message string, partition int32) error
	Close()
}
