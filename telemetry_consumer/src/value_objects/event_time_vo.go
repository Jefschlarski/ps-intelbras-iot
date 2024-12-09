package valueobjects

import (
	"errors"
	"time"
)

// EventTime representa a data do evento
type EventTime struct {
	Seconds int64 `json:"seconds"`
	Nanos   int64 `json:"nanos"`
}

// Validate valida o EventTime
func (e EventTime) Validate() error {
	if e.Seconds >= 0 && e.Nanos >= 0 {
		return nil
	}
	return errors.New("invalid event time")
}

// ToDateTime converte o EventTime para time.Time
func (e EventTime) ToDateTime() time.Time {
	return time.Unix(e.Seconds, e.Nanos)
}
