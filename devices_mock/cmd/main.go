package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "github.com/Jefschlarski/ps-intelbras-iot/devices_mock/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	serverAddress = "telemetry_producer:50051"
	numWorkers    = 5
)

type EventType int

const (
	EventTypeVelocity EventType = 1
	EventTypeRPM      EventType = 2
	EventTypeTemp     EventType = 3
	EventTypeFuel     EventType = 4
	EventTypeMileage  EventType = 5
	EventTypeGPS      EventType = 6
	EventTypeLights   EventType = 7
	EventTypeError    EventType = 8
)

var eventTypes = []EventType{
	EventTypeVelocity,
	EventTypeRPM,
	EventTypeTemp,
	EventTypeFuel,
	EventTypeMileage,
	EventTypeGPS,
	EventTypeLights,
	EventTypeError,
}

func main() {

	// Conexão com o servidor gRPC
	creds := insecure.NewCredentials()
	conn, err := grpc.NewClient(serverAddress, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Não foi possível conectar ao servidor: %v", err)
	}
	defer conn.Close()

	client := pb.NewTelemetryServiceClient(conn)

	var wg sync.WaitGroup

	log.Println("Simulação iniciada.")

	// Inicializa dispositivos simulados
	for i := 0; i < numWorkers; i++ {
		deviceID := i + 1

		for _, eventType := range eventTypes {
			wg.Add(3) // Um para cada tipo de envio (único, batch e stream)

			// Envio de eventos únicos
			go func(deviceID int, eventType EventType) {
				defer wg.Done()
				sendEvent(client, deviceID, int(eventType))
			}(deviceID, eventType)

			// Envio de batch de eventos
			go func(deviceID int, eventType EventType) {
				defer wg.Done()
				sendBatchEvents(client, deviceID, int(eventType))
			}(deviceID, eventType)

			// Stream de eventos
			go func(deviceID int, eventType EventType) {
				defer wg.Done()
				streamEvents(client, deviceID, int(eventType))
			}(deviceID, eventType)
		}
	}

	wg.Wait()

	log.Println("Simulação finalizada.")
}

func sendEvent(client pb.TelemetryServiceClient, deviceID int, eventType int) {
	for {
		event := generateRandomEvent(deviceID, eventType)
		_, err := client.Event(context.Background(), event)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar evento único: %v", deviceID, err)
		}
		time.Sleep(time.Second * 5) // Espera
	}
}

func sendBatchEvents(client pb.TelemetryServiceClient, deviceID int, eventType int) {
	for {
		var batch pb.EventBatch
		for i := 0; i < rand.Intn(20)+1; i++ { // Entre 1 e 20 eventos
			batch.EventBatch = append(batch.EventBatch, generateRandomEvent(deviceID, eventType))
		}
		_, err := client.BatchEvents(context.Background(), &batch)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar batch de eventos: %v", deviceID, err)
		}
		time.Sleep(time.Second * 30) // Espera
	}
}

func streamEvents(client pb.TelemetryServiceClient, deviceID int, eventType int) {
	stream, err := client.StreamEvents(context.Background())
	if err != nil {
		log.Printf("Dispositivo %d: erro ao iniciar stream: %v", deviceID, err)
		return
	}

	for {
		event := generateRandomEvent(deviceID, eventType)
		err := stream.Send(event)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar evento no stream: %v", deviceID, err)
			return
		}
		time.Sleep(time.Second * 10) // Espera
	}
}

func generateRandomEvent(deviceID int, eventType int) *pb.Event {
	switch eventType {
	case int(EventTypeVelocity):
		return generateFloatValueEvent(deviceID, eventType, 0, 200)
	case int(EventTypeRPM):
		return generateIntValueEvent(deviceID, eventType, 0, 8000)
	case int(EventTypeTemp):
		return generateFloatValueEvent(deviceID, eventType, -40, 120)
	case int(EventTypeFuel):
		return generateFloatValueEvent(deviceID, eventType, 0, 100)
	case int(EventTypeMileage):
		return generateIntValueEvent(deviceID, eventType, 0, 1000000)
	case int(EventTypeGPS):
		return generateStringValueEvent(deviceID, eventType)
	case int(EventTypeLights):
		return generateIntValueEvent(deviceID, eventType, 0, 1)
	case int(EventTypeError):
		return generateStringValueEvent(deviceID, eventType)
	default:
		log.Printf("Dispositivo %d: tipo de evento desconhecido %d", deviceID, eventType)
		return nil
	}
}

// Funções auxiliares para gerar eventos
func generateIntValueEvent(deviceID int, eventType int, min, max int) *pb.Event {
	value := rand.Intn(max-min+1) + min
	return &pb.Event{
		DeviceId: int64(deviceID),
		Type:     int32(eventType),
		Time:     timestamppb.New(time.Now()),
		Sensor:   &pb.Event_ValueInt{ValueInt: int32(value)},
	}
}

func generateFloatValueEvent(deviceID int, eventType int, min, max float32) *pb.Event {
	value := min + rand.Float32()*(max-min)
	return &pb.Event{
		DeviceId: int64(deviceID),
		Type:     int32(eventType),
		Time:     timestamppb.New(time.Now()),
		Sensor:   &pb.Event_ValueFloat{ValueFloat: value},
	}
}

func generateStringValueEvent(deviceID int, eventType int) *pb.Event {
	value := "unknown"
	if eventType == int(EventTypeGPS) {
		latitude := fmt.Sprintf("%.6f", -90+rand.Float64()*180)
		longitude := fmt.Sprintf("%.6f", -180+rand.Float64()*360)
		value = fmt.Sprintf("%s, %s", latitude, longitude)
	} else {
		value = errorTypes[rand.Intn(len(errorTypes))]
	}
	return &pb.Event{
		DeviceId: int64(deviceID),
		Type:     int32(eventType),
		Time:     timestamppb.New(time.Now()),
		Sensor:   &pb.Event_ValueString{ValueString: value},
	}
}

var errorTypes = []string{
	"Sensor offline",
	"Overheat detected",
	"Low battery",
	"Data corruption",
	"Invalid value range",
	"Connection timeout",
	"Unauthorized access",
	"Hardware failure",
}
