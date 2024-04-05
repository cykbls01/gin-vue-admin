package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ChartRouter struct{}

func (s *ChartRouter) InitChartRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	chartRouter := Router.Group("chart").Use(middleware.OperationRecord())

	chartRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemChartApi
	{
		chartRouter.POST("createChart", chartRouterApi.CreateChart)
		chartRouter.POST("deleteChart", chartRouterApi.DeleteChart)
		chartRouter.POST("getChartList", chartRouterApi.GetChartList)
		chartRouter.POST("getAllCharts", chartRouterApi.GetAllCharts)
	}
}
