package http

import (
	"github.com/arfanxn/welding/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

func NewRouterFromConfig(cfg *config.Config) *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 10 << 20
	return r
}
