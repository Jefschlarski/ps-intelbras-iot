package repo

import (
	"database/sql"

	"github.com/Jefschlarski/ps-intelbras-iot/telemetry_consumer/src/model"
)

// TelemetryRepo representa o repositorio de telemetria
type TelemetryRepo struct {
	db *sql.DB
}

// NewTelemetryRepo cria um novo repositorio de telemetria
func NewTelemetryRepo(db *sql.DB) TelemetryRepoInterface {
	return &TelemetryRepo{db}
}

// Save salva um registro de telemetria
func (t *TelemetryRepo) Save(telemetry *model.Telemetry) error {
	query, err := t.db.Prepare(`INSERT INTO "telemetry" (device_id, event_type, event_time, value_string, value_float, value_int) VALUES ($1, $2, $3, $4, $5, $6)`)
	if err != nil {
		return err
	}

	defer query.Close()

	_, err = query.Exec(telemetry.Device_id, telemetry.Event_type, telemetry.Event_time, telemetry.Value_string, telemetry.Value_float, telemetry.Value_int)
	if err != nil {
		return err
	}

	defer query.Close()

	return nil
}
