package telemetry

import (
	model "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/model"
	repository "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/repository"
)

type TelemetryService struct {
	r repository.TelemetryRepositoryInterface
}

func NewTelemetryService(repository repository.TelemetryRepositoryInterface) TelemetryServiceInterface {
	return &TelemetryService{r: repository}
}

func (us *TelemetryService) GetTelemetries(limit int) (users []*model.Telemetry, err error) {
	return us.r.GetTelemetries(limit)
}

func (us *TelemetryService) Length() (int, error) {
	return us.r.Length()
}

func (us *TelemetryService) GetDevices() ([]*model.Device, error) {
	return us.r.GetDevices()
}

func (us *TelemetryService) GetTelemetriesByDeviceId(deviceId int64, typeId int64, limit int64) ([]*model.Telemetry, error) {
	return us.r.GetTelemetriesByDeviceId(deviceId, typeId, limit)
}
