package telemetry

import (
	"strconv"

	dto "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/dto"
	service "github.com/Jefschlarski/ps-intelbras-iot/monitor_api/modules/telemetry/service"

	"github.com/gin-gonic/gin"
)

type TelemetryController struct {
	s service.TelemetryServiceInterface
}

func NewTelemetryController(service service.TelemetryServiceInterface) TelemetryControllerInterface {
	return &TelemetryController{s: service}
}

func (uc *TelemetryController) GetTelemetries(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	tekenetries, err := uc.s.GetTelemetries(limit)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if len(tekenetries) == 0 {
		ctx.JSON(204, nil)
		return
	}

	returnTelemetriesDto := []dto.ReturnTelemetryDto{}

	for _, telemetry := range tekenetries {
		returnTelemetriesDto = append(returnTelemetriesDto, dto.ReturnTelemetryDto{
			Device_id:       telemetry.Device_id,
			Event_type:      telemetry.Event_type.GetValue(),
			Event_type_name: telemetry.Event_type.String(),
			Event_time:      telemetry.Event_time.Format("2006-01-02 15:04:05"),
			Value_string:    telemetry.Value_string,
			Value_float:     telemetry.Value_float,
			Value_int:       telemetry.Value_int,
		})
	}

	ctx.JSON(200, returnTelemetriesDto)
}

func (uc *TelemetryController) Length(ctx *gin.Context) {
	len, err := uc.s.Length()
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	ctx.JSON(200, len)
}

func (uc *TelemetryController) GetDevices(ctx *gin.Context) {
	telemetries, err := uc.s.GetDevices()
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	returnTelemetriesDevicesDto := []dto.ReturnTelemetryDevicesDto{}

	for _, telemetry := range telemetries {
		returnTelemetriesDevicesDto = append(returnTelemetriesDevicesDto, dto.ReturnTelemetryDevicesDto{
			Device_id:           telemetry.Device_id,
			Last_telemetry_time: telemetry.Last_telemetry_time.Format("2006-01-02 15:04:05"),
			Last_telemetry_type: telemetry.Last_telemetry_type.String(),
		})
	}

	ctx.JSON(200, telemetries)
}

func (uc *TelemetryController) GetTelemetriesByDeviceId(ctx *gin.Context) {
	deviceId, err := strconv.ParseInt(ctx.Param("device_id"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	typeId, err := strconv.ParseInt(ctx.Request.URL.Query().Get("type"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}
	limit, err := strconv.ParseInt(ctx.Request.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	telemetries, err := uc.s.GetTelemetriesByDeviceId(deviceId, typeId, limit)
	if err != nil {
		ctx.JSON(500, err)
		return
	}

	if len(telemetries) == 0 {
		ctx.JSON(204, nil)
		return
	}

	returnTelemetriesDto := []dto.ReturnTelemetryDto{}

	for _, telemetry := range telemetries {
		returnTelemetriesDto = append(returnTelemetriesDto, dto.ReturnTelemetryDto{
			Device_id:       telemetry.Device_id,
			Event_type:      telemetry.Event_type.GetValue(),
			Event_type_name: telemetry.Event_type.String(),
			Event_time:      telemetry.Event_time.Format("2006-01-02 15:04:05"),
			Value_string:    telemetry.Value_string,
			Value_float:     telemetry.Value_float,
			Value_int:       telemetry.Value_int,
		})
	}

	ctx.JSON(200, returnTelemetriesDto)
}
