package grpc

import (
	context "context"
	"encoding/json"
	"fmt"

	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/dto"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/kafka"
	"github.com/Jefschlarski/ps-intelbras-iot/producer/src/utils"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type grpcService struct {
	producer  kafka.ProducerInterface
	workers   int
	messageCh chan string
	errorCh   chan error
	UnimplementedTelemetryServiceServer
	logger *utils.Logger
}

func NewGrpcService(producer kafka.ProducerInterface, workers int) *grpcService {
	service := &grpcService{
		producer:  producer,
		workers:   workers,
		messageCh: make(chan string),
		errorCh:   make(chan error),
	}

	service.logger = utils.GetLoggerInstance()

	for i := 0; i < workers; i++ {
		go service.worker(0)
	}

	return service
}

func (s *grpcService) worker(partition int32) {
	for message := range s.messageCh {
		if err := s.producer.Produce(message, partition); err != nil {
			s.logger.Error(fmt.Sprintf("Erro ao enviar mensagem: %v", err))
			s.errorCh <- err
		}
	}
}

// SendEvent implementa a RPC SendEvent
func (s *grpcService) Event(ctx context.Context, event *Event) (*emptypb.Empty, error) {
	dto := dto.NewTelemetryEventDto(event.DeviceId, event.Type, event.Time, event.Sensor)

	if err := dto.Validate(); err != nil {
		s.logger.Error(fmt.Sprintf("Dados inválidos: %v", err))
		return &emptypb.Empty{}, err
	}

	jsonData, err := json.Marshal(dto)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Erro ao converter para JSON: %v", err))
		return &emptypb.Empty{}, err
	}

	select {
	case s.messageCh <- string(jsonData):
	case <-ctx.Done():
		return &emptypb.Empty{}, ctx.Err()
	}

	select {
	case err := <-s.errorCh:
		return &emptypb.Empty{}, fmt.Errorf("failed to produce message: %w", err)
	default:
	}

	return &emptypb.Empty{}, nil
}

// BatchEvents implementa a RPC BatchEvents
func (s *grpcService) BatchEvents(ctx context.Context, batch *EventBatch) (*emptypb.Empty, error) {
	for _, event := range batch.EventBatch {
		dto := dto.NewTelemetryEventDto(event.DeviceId, event.Type, event.Time, event.Sensor)
		if err := dto.Validate(); err != nil {
			s.logger.Error(fmt.Sprintf("Dados inválidos: %v", err))
			return &emptypb.Empty{}, err
		}

		jsonData, err := json.Marshal(dto)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Erro ao converter para JSON: %v", err))
			return &emptypb.Empty{}, err
		}

		select {
		case s.messageCh <- string(jsonData):
		case <-ctx.Done():
			return &emptypb.Empty{}, ctx.Err()
		}
	}

	return &emptypb.Empty{}, nil
}

// StreamEvents implementa a RPC StreamEvents
func (s *grpcService) StreamEvents(stream TelemetryService_StreamEventsServer) error {
	for {
		event, err := stream.Recv()
		if err != nil {
			s.logger.Error(fmt.Sprintf("Erro ao receber mensagem: %v", err))
			return err
		}

		dto := dto.NewTelemetryEventDto(event.DeviceId, event.Type, event.Time, event.Sensor)
		if err := dto.Validate(); err != nil {
			s.logger.Error(fmt.Sprintf("Dados inválidos: %v", err))
			return err
		}

		jsonData, err := json.Marshal(dto)
		if err != nil {
			s.logger.Error(fmt.Sprintf("Erro ao converter para JSON: %v", err))
			return err
		}

		select {
		case s.messageCh <- string(jsonData):
		case <-stream.Context().Done():
			return stream.Context().Err()
		}
	}
}

// GracefulStop encerra os workers e aguarda a conclusão deles
func (s *grpcService) GracefulStop() {

	s.logger.Info("Encerrando o grpc service")
	close(s.messageCh)
	for i := 0; i < s.workers; i++ {
		<-s.errorCh
	}
	close(s.errorCh)
}
