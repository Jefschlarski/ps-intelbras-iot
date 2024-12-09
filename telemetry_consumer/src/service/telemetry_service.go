package service

import (
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/dto"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/model"
	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/repo"
)

// TelemetryService representa o serviço de telemetria
type TelemetryService struct {
	telemetryRepo repo.TelemetryRepoInterface
}

// NewTelemetryService cria um novo serviço de telemetria
func NewTelemetryService(telemetryRepo repo.TelemetryRepoInterface) TelemetryServiceInterface {
	return &TelemetryService{telemetryRepo}
}

// Save salva um registro de telemetria
func (s *TelemetryService) Save(telemetry *dto.TelemetryEventDto) error {
	model := model.TelemetryFromDto(telemetry)
	return s.telemetryRepo.Save(model)
}
