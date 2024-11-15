package business

import (
	"github.com/giles-wong/zhongx-gva/server/middleware"
	"github.com/gin-gonic/gin"
)

type CertReadingRouter struct {
}

func (cr *CertReadingRouter) InitCertReadingRouter(Router *gin.RouterGroup) {
	readingRouter := Router.Group("cert").Use(middleware.OperationRecord())
	readingRouterWithoutRecord := Router.Group("cert")
	{
		readingRouter.POST("addReading", certReadingApi.AddReading)         // 管理员注册账号
		readingRouter.POST("editReading", certReadingApi.EditReading)       // 用户修改密码
		readingRouter.DELETE("deleteReading", certReadingApi.DeleteReading) // 用户修改密码
	}
	{
		readingRouterWithoutRecord.POST("getReadingList", certReadingApi.GetReadingList) // 分页获取证书列表
		//readingRouterWithoutRecord.GET("getInfo", certReadingApi.GetReadingInfo)  // 获取证书信息
	}
}
