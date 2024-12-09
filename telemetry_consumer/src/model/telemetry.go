package model

import (
	"strconv"
	"time"

	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/dto"
	valueobjects "github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/value_objects"
)

// Telemetry representa um registro de telemetria
type Telemetry struct {
	Device_id    int64                  `json:"device_id"`
	Event_type   valueobjects.EventType `json:"event_type"`
	Event_time   time.Time              `json:"event_time"`
	Value_string string                 `json:"value_string"`
	Value_float  float64                `json:"value_float"`
	Value_int    int64                  `json:"value_int"`
}

// NewTelemetry cria um novo registro de telemetria
func NewTelemetry(device_id int64, event_type int32, event_time time.Time, value_string string, value_float float64, value_int int64) *Telemetry {
	return &Telemetry{
		Device_id:    device_id,
		Event_type:   valueobjects.EventType(event_type),
		Event_time:   event_time,
		Value_string: value_string,
		Value_float:  value_float,
		Value_int:    value_int,
	}
}

// TelemetryFromDto cria um novo registro de telemetria a partir de um dto TelemetryEventDto
func TelemetryFromDto(dto *dto.TelemetryEventDto) *Telemetry {

	return &Telemetry{
		Device_id:    dto.Device_id,
		Event_type:   dto.Event_type,
		Event_time:   dto.Event_time.ToDateTime(),
		Value_string: dto.Sensor_value.Value_string,
		Value_float:  dto.Sensor_value.Value_float,
		Value_int:    dto.Sensor_value.Value_int,
	}
}

// ToString converte o Telemetry para string
func (t *Telemetry) ToString() string {
	return "Telemetry{" +
		"Device_id=" + strconv.FormatInt(t.Device_id, 10) +
		", Event_type=" + t.Event_type.String() +
		", Event_time=" + t.Event_time.String() +
		", Value_string=" + t.Value_string +
		", Value_float=" + strconv.FormatFloat(t.Value_float, 'f', -1, 64) +
		", Value_int=" + strconv.FormatInt(t.Value_int, 10) +
		"}"
}
