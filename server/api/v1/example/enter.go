package example

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CustomerApi
	FileUploadAndDownloadApi
}

var (
	customerService              = service.ServiceGroupApp.ExampleServiceGroup.CustomerService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
	clusterService               = service.ServiceGroupApp.SystemServiceGroup.ClusterService
	chartService                 = service.ServiceGroupApp.SystemServiceGroup.ChartService
)

func ttt() {
	chartService.GetAllCharts()
}
