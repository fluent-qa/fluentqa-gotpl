package api

import (
	"github.com/fluent-qa/gobase/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func NewServerHTTPWithMiddleware() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(
		middleware.CORSMiddleware(),
	)
	return r
}
