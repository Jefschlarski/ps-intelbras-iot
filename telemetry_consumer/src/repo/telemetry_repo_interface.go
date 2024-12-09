package repo

import "github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/model"

type TelemetryRepoInterface interface {
	Save(telemetry *model.Telemetry) error
}
