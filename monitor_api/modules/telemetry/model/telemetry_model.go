package telemetry

import (
	"time"

	valueobjects "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/value_objects"
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

type Device struct {
	Device_id           int64                  `json:"device_id"`
	Last_telemetry_time time.Time              `json:"last_telemetry_time"`
	Last_telemetry_type valueobjects.EventType `json:"last_telemetry_type"`
}
