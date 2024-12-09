package valueobjects

import "errors"

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

func (e EventType) Validate() error {
	if e >= EventTypeVelocity && e <= EventTypeError {
		return nil
	}
	return errors.New("invalid event type")
}
