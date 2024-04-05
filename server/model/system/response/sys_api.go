package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysAPIResponse struct {
	Api system.SysApi `json:"api"`
}

type SysAPIListResponse struct {
	Apis []system.SysApi `json:"apis"`
}

type SysClusterListResponse struct {
	Clusters []system.SysCluster `json:"clusters"`
}

type SysChartListResponse struct {
	Charts []system.SysChart `json:"charts"`
}
