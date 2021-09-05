package prometheus_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusHandler() gin.HandlerFunc {
	promHandler := promhttp.Handler()
	return func(g *gin.Context) {
		promHandler.ServeHTTP(g.Writer, g.Request)
	}
}
