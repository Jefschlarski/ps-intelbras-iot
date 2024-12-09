package service

import (
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/dto"
)

// TelemetryServiceInterface representa a interface do serviço de telemetria
type TelemetryServiceInterface interface {
	Save(telemetry *dto.TelemetryEventDto) error
}
