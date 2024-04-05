package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SystemComponentApi struct{}

func (s *SystemComponentApi) CreateComponent(c *gin.Context) {
	var component system.SysComponent
	err := c.ShouldBindJSON(&component)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	component.Status = "创建中"
	err = componentService.CreateComponent(&component)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	chart, err := chartService.GetChartById(int(component.ChartID))
	if err != nil {
		global.GVA_LOG.Error("获取chart失败!", zap.Error(err))
		response.FailWithMessage("获取chart失败", c)
		return
	}

	cluster, err := clusterService.GetClusterById(int(component.ClusterID))
	if err != nil {
		global.GVA_LOG.Error("获取cluster失败!", zap.Error(err))
		response.FailWithMessage("获取cluster失败", c)
		return
	}

	go func() {
		args := system.SysHelmRpc{
			Cluster:   cluster,
			Component: component,
			Chart:     chart,
		}
		err, _ = utils.Call("HelmRpc", "Create", args)
		if err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			return
		}

		err, component = utils.Call("HelmRpc", "Port", args)
		if err != nil {
			global.GVA_LOG.Error("获取端口失败!", zap.Error(err))
			return
		}
		err = componentService.UpdateComponent(component)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			return
		}
	}()
	response.OkWithMessage("创建成功", c)
}

func (s *SystemComponentApi) DeleteComponent(c *gin.Context) {
	var component system.SysComponent
	err := c.ShouldBindJSON(&component)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	chart, err := chartService.GetChartById(int(component.ChartID))
	if err != nil {
		global.GVA_LOG.Error("获取chart失败!", zap.Error(err))
		response.FailWithMessage("获取chart失败", c)
		return
	}

	cluster, err := clusterService.GetClusterById(int(component.ClusterID))
	if err != nil {
		global.GVA_LOG.Error("获取cluster失败!", zap.Error(err))
		response.FailWithMessage("获取cluster失败", c)
		return
	}
	args := system.SysHelmRpc{
		Cluster:   cluster,
		Component: component,
		Chart:     chart,
	}

	go func() {
		err, component = utils.Call("HelmRpc", "Delete", args)
		if err != nil {
			global.GVA_LOG.Error("删除失败!", zap.Error(err))
			return
		}
		
		err = componentService.UpdateComponent(component)
		if err != nil {
			global.GVA_LOG.Error("更新失败!", zap.Error(err))
			return
		}
	}()

	response.OkWithMessage("删除成功", c)
}

func (s *SystemComponentApi) GetComponentList(c *gin.Context) {
	var pageInfo systemReq.SearchApiParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := componentService.GetComponentInfoList(pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (s *SystemComponentApi) GetGrafanaLink(c *gin.Context) {
	type res struct {
		Link string `json:"link"`
	}
	response.OkWithDetailed(res{Link: global.GVA_CONFIG.Helm.Grafana}, "获取成功", c)
}

func (s *SystemComponentApi) GetHelmConfig(c *gin.Context) {
	type res struct {
		Data config.Helm `json:"data"`
	}
	response.OkWithDetailed(res{Data: global.GVA_CONFIG.Helm}, "获取成功", c)
}

func (s *SystemComponentApi) HandleDownloadFile(c *gin.Context) {
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "inline; filename="+"tutorial.pdf")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File("uploads/file/tutorial.pdf")
}
