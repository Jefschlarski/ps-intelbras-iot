package telemetry

import (
	"database/sql"

	model "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/model"
)

type TelemetryRepository struct {
	db *sql.DB
}

func NewTelemetryRepository(db *sql.DB) TelemetryRepositoryInterface {
	return &TelemetryRepository{db: db}
}

func (tr *TelemetryRepository) GetTelemetries(limit int) (telemetries []*model.Telemetry, err error) {
	query, err := tr.db.Prepare(`SELECT device_id, event_type, event_time, value_string, value_float, value_int FROM "telemetry" order by event_time desc limit $1`)
	if err != nil {
		return nil, err
	}

	defer query.Close()

	rows, err := query.Query(limit)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		telemetry := &model.Telemetry{}
		err = rows.Scan(
			&telemetry.Device_id,
			&telemetry.Event_type,
			&telemetry.Event_time,
			&telemetry.Value_string,
			&telemetry.Value_float,
			&telemetry.Value_int)
		if err != nil {
			return nil, err
		}

		telemetries = append(telemetries, telemetry)
	}

	return telemetries, nil
}

func (tr *TelemetryRepository) Length() (int, error) {
	var len int
	err := tr.db.QueryRow(`SELECT count(*) FROM "telemetry"`).Scan(&len)
	return len, err
}

func (tr *TelemetryRepository) GetDevices() ([]*model.Device, error) {
	var devices []*model.Device

	query, err := tr.db.Prepare(`SELECT device_id, event_time, event_type
FROM (
    SELECT device_id, event_time, event_type,
           ROW_NUMBER() OVER (PARTITION BY device_id ORDER BY event_time DESC) AS rn
    FROM "telemetry"
) AS ranked
WHERE rn = 1
ORDER BY device_id DESC;`)
	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		device := &model.Device{}
		err = rows.Scan(&device.Device_id, &device.Last_telemetry_time, &device.Last_telemetry_type)
		if err != nil {
			return nil, err
		}
		devices = append(devices, device)
	}
	return devices, err
}

func (tr *TelemetryRepository) GetTelemetriesByDeviceId(deviceId int64, typeId int64, limit int64) ([]*model.Telemetry, error) {
	query, err := tr.db.Prepare(`SELECT device_id, event_type, event_time, value_string, value_float, value_int FROM "telemetry" where device_id = $1 and event_type = $2 order by event_time asc limit $3`)

	if err != nil {
		return nil, err
	}
	defer query.Close()

	rows, err := query.Query(deviceId, typeId, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var telemetries []*model.Telemetry
	for rows.Next() {
		telemetry := &model.Telemetry{}
		err = rows.Scan(
			&telemetry.Device_id,
			&telemetry.Event_type,
			&telemetry.Event_time,
			&telemetry.Value_string,
			&telemetry.Value_float,
			&telemetry.Value_int)
		if err != nil {
			return nil, err
		}
		telemetries = append(telemetries, telemetry)
	}
	return telemetries, err
}
