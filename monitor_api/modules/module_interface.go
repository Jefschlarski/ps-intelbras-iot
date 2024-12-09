package modules

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type ModuleInterface interface {
	Init(router *gin.RouterGroup, db *sql.DB)
}
