package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/logging"
	"github.com/gin-gonic/gin"
)

func StructuredLog(logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.FullPath()
		raw := ctx.Request.URL.RawQuery
		reqBdy, _ := io.ReadAll(ctx.Request.Body)
		ctx.Request.Body.Close()
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBdy))

		ctx.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.BodySize = ctx.Writer.Size()

		if raw != "" {
			path += "?" + raw
		}
		param.Path = path

		key := map[logging.ExtraKey]interface{}{}

		key[logging.ClientIp] = param.ClientIP
		key[logging.Method] = param.Method
		key[logging.Latency] = param.Latency
		key[logging.BodySize] = param.BodySize
		key[logging.Path] = param.Path
		key[logging.RequestBody] = string(reqBdy)

		logger.Info(logging.RequestResponse, logging.Api, "", key)

	}
}
