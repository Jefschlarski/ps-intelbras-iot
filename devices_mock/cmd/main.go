package main

import (
	"context"
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
	serverAddress = "localhost:50051"
	numDevices    = 20
)

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

	// Inicializa dispositivos simulados
	for i := 0; i < numDevices; i++ {
		wg.Add(3)
		deviceID := i + 1

		// Envio de eventos únicos
		go func() {
			defer wg.Done()
			sendEvent(client, deviceID)
		}()

		// Envio de batch de eventos
		go func() {
			defer wg.Done()
			sendBatchEvents(client, deviceID)
		}()

		// Stream de eventos
		go func() {
			defer wg.Done()
			streamEvents(client, deviceID)
		}()
	}

	wg.Wait()
	log.Println("Simulação finalizada.")
}

func sendEvent(client pb.TelemetryServiceClient, deviceID int) {
	for {
		event := generateRandomEvent(deviceID)
		_, err := client.SendEvent(context.Background(), event)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar evento único: %v", deviceID, err)
		} else {
			log.Printf("Dispositivo %d: evento único enviado com sucesso", deviceID)
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(5)+1)) // Espera aleatória
	}
}

func sendBatchEvents(client pb.TelemetryServiceClient, deviceID int) {
	for {
		var batch pb.EventBatch
		for i := 0; i < rand.Intn(10)+1; i++ { // Entre 1 e 10 eventos
			batch.EventBatch = append(batch.EventBatch, generateRandomEvent(deviceID))
		}
		_, err := client.BatchEvents(context.Background(), &batch)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar batch de eventos: %v", deviceID, err)
		} else {
			log.Printf("Dispositivo %d: batch de eventos enviado com sucesso", deviceID)
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(10)+5)) // Espera aleatória
	}
}

func streamEvents(client pb.TelemetryServiceClient, deviceID int) {
	stream, err := client.StreamEvents(context.Background())
	if err != nil {
		log.Printf("Dispositivo %d: erro ao iniciar stream: %v", deviceID, err)
		return
	}

	for {
		event := generateRandomEvent(deviceID)
		err := stream.Send(event)
		if err != nil {
			log.Printf("Dispositivo %d: erro ao enviar evento no stream: %v", deviceID, err)
			return
		}
		log.Printf("Dispositivo %d: evento enviado no stream", deviceID)
		time.Sleep(time.Second * time.Duration(rand.Intn(3)+1)) // Espera aleatória
	}
}

func generateRandomEvent(deviceID int) *pb.Event {
	// Gera dados aleatórios para o evento
	sensorType := rand.Intn(3) // 0: int, 1: float, 2: string
	var event *pb.Event
	switch sensorType {
	case 0:
		event = generateIntValueEvent(deviceID)
	case 1:
		event = generateFloatValueEvent(deviceID)
	case 2:
		event = generateStringValueEvent(deviceID)
	}

	return event
}

func generateIntValueEvent(deviceID int) *pb.Event {
	sensorValue := &pb.Event_ValueInt{ValueInt: rand.Int31n(100)}

	return &pb.Event{
		Type:   int32(deviceID),
		Time:   timestamppb.New(time.Now()),
		Sensor: sensorValue,
	}
}

func generateFloatValueEvent(deviceID int) *pb.Event {
	sensorValue := &pb.Event_ValueFloat{ValueFloat: rand.Float32() * 100}

	return &pb.Event{
		Type:   int32(deviceID),
		Time:   timestamppb.New(time.Now()),
		Sensor: sensorValue,
	}
}

func generateStringValueEvent(deviceID int) *pb.Event {
	sensorValue := &pb.Event_ValueString{ValueString: randomString(10)}

	return &pb.Event{
		Type:   int32(deviceID),
		Time:   timestamppb.New(time.Now()),
		Sensor: sensorValue,
	}
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randStr := make([]rune, length)
	for i := range randStr {
		randStr[i] = letters[rand.Intn(len(letters))]
	}
	return string(randStr)
}
