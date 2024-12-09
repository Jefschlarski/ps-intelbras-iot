package kafka

// ConsumerInterface representa a interface que deve ser implementada por um consumidor de mensagens
type ConsumerInterface interface {
	Consume() error
	Close()
}
