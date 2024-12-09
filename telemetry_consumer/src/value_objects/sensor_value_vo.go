package valueobjects

import (
	"errors"
)

// SensorValue representa o valor do sensor
type SensorValue struct {
	Value_string string  `json:"ValueString"`
	Value_float  float64 `json:"ValueFloat"`
	Value_int    int64   `json:"ValueInt"`
}

// Validate valida o SensorValue
func (s SensorValue) Validate() error {
	if s.Value_string != "" || s.Value_float != 0 || s.Value_int != 0 {
		return nil
	}
	return errors.New("invalid sensor value")
}
