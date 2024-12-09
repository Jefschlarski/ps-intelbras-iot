package dto

import (
	"errors"

	valueobjects "github.com/Jefschlarski/ps-intelbras-iot/producer/src/value_objects"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// -Velocidade do veículo

// -RPM do motor

// -Temperatura do motor

// -Nível de combustível

// -Quilometragem percorrida

// -Localização GPS

// -Status das luzes (faróis, lanternas, etc.)

type TelemetryEventDto struct {
	Device_id    int64                  `json:"device_id"`
	Event_type   valueobjects.EventType `json:"event_type"`
	Event_time   *timestamppb.Timestamp `json:"event_time"`
	Sensor_value interface{}            `json:"sensor_value"`
}

func NewTelemetryEventDto(device_id int64, event_type int32, event_time *timestamppb.Timestamp, sensor_value interface{}) *TelemetryEventDto {
	return &TelemetryEventDto{
		Device_id:    device_id,
		Event_type:   valueobjects.EventType(event_type),
		Event_time:   event_time,
		Sensor_value: sensor_value,
	}
}

func (t *TelemetryEventDto) Validate() error {
	if t.Device_id <= 0 {
		return errors.New("device_id must be greater than 0")
	}

	err := t.Event_type.Validate()
	if err != nil {
		return err
	}

	if t.Event_time == nil {
		return errors.New("event_time must not be nil")
	}
	if t.Sensor_value == nil {
		return errors.New("sensor_value must not be nil")
	}
	return nil
}
