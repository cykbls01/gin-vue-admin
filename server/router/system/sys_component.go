package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComponentRouter struct{}

func (s *ComponentRouter) InitComponentRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	componentRouter := Router.Group("component").Use(middleware.OperationRecord())

	componentRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemComponentApi
	{
		componentRouter.POST("createComponent", componentRouterApi.CreateComponent)
		componentRouter.POST("deleteComponent", componentRouterApi.DeleteComponent)
		componentRouter.POST("getComponentList", componentRouterApi.GetComponentList)
		componentRouter.POST("getGrafanaLink", componentRouterApi.GetGrafanaLink)
		componentRouter.POST("getHelmConfig", componentRouterApi.GetHelmConfig)
		componentRouter.GET("download", componentRouterApi.HandleDownloadFile)
	}
}
