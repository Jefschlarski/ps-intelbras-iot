package main

import (
	"context"
	"log"
	"net"

	pb "github.com/Jefschlarski/ps-intelbras-iot/producer/src/grpc"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/kafka"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server é a implementação do serviço gRPC definido no arquivo proto
type server struct {
	kafka kafka.KafkaProducer
	pb.UnimplementedTelemetryServiceServer
}

func NewServer(kafka kafka.KafkaProducer) *server {
	return &server{kafka: kafka}
}

// SendEvent implementa a RPC SendEvent
func (s *server) SendEvent(ctx context.Context, event *pb.Event) (*pb.Response, error) {
	message, err := protobufToJSON(event)
	if err != nil {
		return nil, err
	}
	s.kafka.Produce(message)

	log.Printf("Recebido evento: type=%v time=%v", event.Type, event.Time.AsTime())
	switch v := event.Sensor.(type) {
	case *pb.Event_ValueInt:
		log.Printf("Valor int: %d", v.ValueInt)
	case *pb.Event_ValueFloat:
		log.Printf("Valor float: %f", v.ValueFloat)
	case *pb.Event_ValueString:
		log.Printf("Valor string: %s", v.ValueString)
	}
	return &pb.Response{EmptyField: &emptypb.Empty{}}, nil
}
func protobufToJSON(msg proto.Message) (string, error) {
	jsonString, err := protojson.Marshal(msg)
	return string(jsonString), err
}

// BatchEvents implementa a RPC BatchEvents
func (s *server) BatchEvents(ctx context.Context, batch *pb.EventBatch) (*pb.Response, error) {
	log.Printf("Recebido batch de eventos, total=%d", len(batch.EventBatch))
	for _, event := range batch.EventBatch {
		log.Printf("Processando evento: type=%v", event.Type)
	}
	return &pb.Response{EmptyField: &emptypb.Empty{}}, nil
}

// StreamEvents implementa a RPC StreamEvents
func (s *server) StreamEvents(stream pb.TelemetryService_StreamEventsServer) error {
	log.Println("StreamEvents iniciado")
	for {
		event, err := stream.Recv()
		if err != nil {
			log.Printf("Erro ao receber stream: %v", err)
			return err
		}
		log.Printf("Evento recebido no stream: type=%v", event.Type)
	}
}

func main() {
	brokers := []string{"localhost:9092"}
	topic := "telemetry"
	kafka, err := kafka.NewKafkaProducer(brokers, topic)
	if err != nil {
		log.Println(err)
		log.Fatalf("Falha ao iniciar a conexão com os brokers kafka")
	}

	// Inicializa o listener na porta 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao iniciar o listener: %v", err)
	}

	// Cria um servidor gRPC
	grpcServer := grpc.NewServer()

	telemetry := NewServer(*kafka)

	// Registra o serviço TelemetryService
	pb.RegisterTelemetryServiceServer(grpcServer, telemetry)

	log.Println("Servidor gRPC iniciado na porta 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}
