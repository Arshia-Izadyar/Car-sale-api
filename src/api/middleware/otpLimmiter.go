package middleware

import (
	"errors"
	"net/http"
	"time"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/Arshia-Izadyar/Car-sale-api/src/config"
	"github.com/Arshia-Izadyar/Car-sale-api/src/pkg/limiter"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIpRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(ctx *gin.Context) {
		limiter := limiter.GetLimiter(ctx.ClientIP())
		if !limiter.Allow() {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, int(helper.OtpLimiterError), errors.New("requested too many otp codes").Error()))
			return
		}
		ctx.Next()
	}
}
