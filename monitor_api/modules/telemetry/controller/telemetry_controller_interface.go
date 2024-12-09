package telemetry

import (
	"github.com/gin-gonic/gin"
)

type TelemetryControllerInterface interface {
	GetTelemetries(ctx *gin.Context)
	Length(ctx *gin.Context)
	GetDevices(ctx *gin.Context)
	GetTelemetriesByDeviceId(ctx *gin.Context)
}
