package telemetry

type ReturnTelemetryDto struct {
	Device_id       int64   `json:"device_id"`
	Event_type      int32   `json:"event_type"`
	Event_type_name string  `json:"event_type_name"`
	Event_time      string  `json:"event_time"`
	Value_string    string  `json:"value_string"`
	Value_float     float64 `json:"value_float"`
	Value_int       int64   `json:"value_int"`
}

type ReturnTelemetryDevicesDto struct {
	Device_id           int64  `json:"device_id"`
	Last_telemetry_time string `json:"last_telemetry_time"`
	Last_telemetry_type string `json:"last_telemetry_type"`
}
