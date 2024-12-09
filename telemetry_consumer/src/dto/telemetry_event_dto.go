package dto

import (
	"errors"

	valueobjects "github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/value_objects"
)

// TelemetryEventDto representa um dto de telemetria
type TelemetryEventDto struct {
	Device_id    int64                     `json:"device_id"`
	Event_type   valueobjects.EventType    `json:"event_type"`
	Event_time   valueobjects.EventTime    `json:"event_time"`
	Sensor_value *valueobjects.SensorValue `json:"sensor_value"`
}

// Validate valida o TelemetryEventDto
func (t *TelemetryEventDto) Validate() error {
	if t.Device_id <= 0 {
		return errors.New("device_id must be greater than 0")
	}

	err := t.Event_type.Validate()
	if err != nil {
		return err
	}

	err = t.Event_time.Validate()
	if err != nil {
		return err
	}

	err = t.Sensor_value.Validate()
	if err != nil {
		return err
	}

	return nil
}
