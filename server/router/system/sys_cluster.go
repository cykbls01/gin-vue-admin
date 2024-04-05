package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ClusterRouter struct{}

func (s *ClusterRouter) InitClusterRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	clusterRouter := Router.Group("cluster").Use(middleware.OperationRecord())

	clusterRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemClusterApi
	{
		clusterRouter.POST("createCluster", clusterRouterApi.CreateCluster)   // 创建Api
		clusterRouter.POST("deleteCluster", clusterRouterApi.DeleteCluster)   // 创建Api
		clusterRouter.POST("getClusterList", clusterRouterApi.GetClusterList) // 创建Api
		clusterRouter.POST("getAllClusters", clusterRouterApi.GetAllClusters)
	}
}
