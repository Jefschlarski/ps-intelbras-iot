package telemetry

import (
	"database/sql"

	"github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules"
	controller "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/controller"
	repository "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/repository"
	service "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/service"
	m "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

type TelemetryModule struct{}

func NewTelemetryModule() modules.ModuleInterface {
	return &TelemetryModule{}
}

func (um *TelemetryModule) Init(router *gin.RouterGroup, db *sql.DB) {
	telemetryRepository := repository.NewTelemetryRepository(db)
	telemetryService := service.NewTelemetryService(telemetryRepository)
	telemetryController := controller.NewTelemetryController(telemetryService)

	router.GET("/telemetries/:limit", m.Logger(telemetryController.GetTelemetries))
	router.GET("/telemetries/length", m.Logger(telemetryController.Length))
	router.GET("/telemetries/devices", m.Logger(telemetryController.GetDevices))
	router.GET("/telemetries/devices/:device_id", m.Logger(telemetryController.GetTelemetriesByDeviceId))
}
