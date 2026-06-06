package middlewares

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kishanknows/Ecom-backend/internal/metrics"
)

func MetricsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		metrics.HttpRequestsTotal.WithLabelValues(
			ctx.Request.Method,
			ctx.FullPath(),
			strconv.Itoa(ctx.Writer.Status()),
		).Inc()

		metrics.HttpRequestDuration.WithLabelValues(
			ctx.Request.Method,
			ctx.FullPath(),
		).Observe(time.Since(start).Seconds())
	}
}