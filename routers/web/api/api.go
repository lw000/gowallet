package api

import (
	"fmt"
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// RegisterService 注册服务
func RegisterService(engine *gin.Engine) {
	api := engine.Group("/wallet")

	lmt := tollbooth.NewLimiter(200, &limiter.ExpirableOptions{ExpireJobInterval: time.Second})
	lmt.SetMessage(fmt.Sprintf("{\"code\":-1,\"msg\":\"too many requests\"}"))

	api.GET("/api", tollbooth_gin.LimitHandler(lmt), channelHandler)
}

func channelHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"c": 0, "m": "what do you want to do?", "d": gin.H{}})
}
