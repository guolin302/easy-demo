package router

import (
	"easy_demo/service"
	"github.com/gin-gonic/gin"
)

func Initrouter() *gin.Engine {
	r := gin.Default()

	pulicGroup := r.Group("/api")
	{
		ownerRouter := pulicGroup.Group("/job")
		{
			ownerRouter.GET("/content/work", service.EchoDomain)
		}
	}
	return r
}
