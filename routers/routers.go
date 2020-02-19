package routers

import (
	"github.com/gin-gonic/gin"
	"gowallet/routers/web/api"
)

func RegisterService(engine *gin.Engine) {
	// 注册子服务
	api.RegisterService(engine)
}
