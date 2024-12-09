package telemetry

import (
	model "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/model"
)

type TelemetryRepositoryInterface interface {
	GetTelemetries(limit int) (users []*model.Telemetry, err error)
	Length() (int, error)
	GetDevices() ([]*model.Device, error)
	GetTelemetriesByDeviceId(deviceId int64, typeId int64, limit int64) ([]*model.Telemetry, error)
}
