package system

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type SysChart struct {
	global.GVA_MODEL
	Name        string `json:"name"`
	Path        string `json:"path"`
	Description string `json:"description"`
	Config      string `json:"config"`
	Version     string `json:"version"`
	Category    string `json:"category"`
}

type SysCluster struct {
	global.GVA_MODEL
	IpGroup     string `json:"ipGroup" gorm:"comment:ip组"`
	Path        string `json:"path" gorm:"comment:配置文件路径"`
	Name        string `json:"name" gorm:"comment:集群名称"`
	Description string `json:"description"`
	Config      string `json:"config" gorm:"comment:k8s配置文件"`
}

type SysComponent struct {
	global.GVA_MODEL
	Name        string `json:"name"`
	Namespace   string `json:"namespace"`
	ChartID     uint   `json:"chartID"`
	Description string `json:"description"`
	ClusterID   uint   `json:"clusterID"`
	Port        string `json:"port"`
	Host        string `json:"host"`
	Password    string `json:"password"`
	Cmd         string `json:"cmd" gorm:"type:text"`
	Status      string `json:"status"`
}

type SysHelmRpc struct {
	Component SysComponent
	Chart     SysChart
	Cluster   SysCluster
}

func (SysCluster) TableName() string {
	return "sys_clusters"
}

func (SysChart) TableName() string {
	return "sys_charts"
}

func (SysComponent) TableName() string {
	return "sys_components"
}
