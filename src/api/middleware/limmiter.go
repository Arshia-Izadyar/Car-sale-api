package middleware

import (
	"errors"
	"net/http"

	"github.com/Arshia-Izadyar/Car-sale-api/src/api/helper"
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func Limiter() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, int(helper.LimiterError), errors.New("you sent too many request please cool down a little").Error()))
			return
		}
		ctx.Next()
	}
}
