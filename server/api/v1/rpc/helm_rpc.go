package rpc

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system/Component"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"reflect"
)

type HelmRpc struct {
}

var chartService = service.ServiceGroupApp.SystemServiceGroup.ChartService
var clusterService = service.ServiceGroupApp.SystemServiceGroup.ClusterService
var componentService = service.ServiceGroupApp.SystemServiceGroup.ComponentService

func AssignStructFields(src interface{}, target interface{}) error {
	// 获取src和target的反射类型对象
	srcType := reflect.TypeOf(src)
	targetType := reflect.TypeOf(target)

	// 确保src是一个结构体
	if srcType.Kind() != reflect.Struct {
		return fmt.Errorf("src must be a struct")
	}

	// 确保target是指向结构体的指针
	if targetType.Kind() != reflect.Ptr || targetType.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("target must be a pointer to a struct")
	}

	// 确保src和target指向的结构体类型相同
	if srcType != targetType.Elem() {
		return fmt.Errorf("src and target must be of the same struct type")
	}

	// 获取src和target的反射值对象
	srcValue := reflect.ValueOf(src)
	targetValue := reflect.ValueOf(target).Elem() // 解引用指针

	// 遍历src的所有字段并赋值给target的对应字段
	for i := 0; i < srcType.NumField(); i++ {
		srcField := srcValue.Field(i)
		targetField := targetValue.Field(i)

		// 如果字段是可设置的（即不是未导出的），则进行赋值
		if targetField.CanSet() {
			targetField.Set(srcField)
		}
	}

	return nil
}

func (rpc *HelmRpc) Create(ctx context.Context, args system.SysHelmRpc, component *system.SysComponent) error {
	chart := args.Chart
	cluster := args.Cluster
	component = &args.Component
	str := fmt.Sprintf("helm install %s -n %s --kubeconfig %s -f %s %s", component.Name, component.Namespace, global.GVA_CONFIG.Helm.Path+cluster.Config, global.GVA_CONFIG.Helm.Path+chart.Config, global.GVA_CONFIG.Helm.Repo+chart.Path)
	if _, err := utils.Exec(str); err != nil {
		return err
	}
	return nil
}

func (rpc *HelmRpc) Delete(ctx context.Context, args system.SysHelmRpc, component *system.SysComponent) error {
	cluster := args.Cluster
	AssignStructFields(args.Component, component)
	str := fmt.Sprintf("helm uninstall %s -n %s --kubeconfig %s", component.Name, component.Namespace, cluster.Config)
	if _, err := utils.Exec(str); err != nil {
		return err
	}
	component.Status = "已回收"
	return nil
}

func (rpc *HelmRpc) Port(ctx context.Context, args system.SysHelmRpc, component *system.SysComponent) error {
	chart := args.Chart
	cluster := args.Cluster
	AssignStructFields(args.Component, component)
	var task = Component.Convert(chart.Category)
	if port, err := utils.Exec(task.Port(cluster, *component)); err != nil {
		return err
	} else {
		component.Port = port
		component.Host = cluster.IpGroup
		component.Status = "运行中"
		return nil
	}
}

func (rpc *HelmRpc) Test(ctx context.Context, args system.SysComponent, component *system.SysComponent) error {
	AssignStructFields(args, component)
	component.Status = "运行中"
	return nil
}
