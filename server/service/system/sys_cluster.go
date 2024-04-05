package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type ClusterService struct{}

var ClusterServiceImpl = new(ClusterService)

func (clusterService *ClusterService) CreateCluster(cluster system.SysCluster) (err error) {
	if !errors.Is(global.GVA_DB.Where("ip_group = ?", cluster.IpGroup).First(&system.SysCluster{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同cluster")
	}
	return global.GVA_DB.Create(&cluster).Error
}

func (clusterService *ClusterService) UpdateCluster(cluster system.SysCluster) (err error) {
	return global.GVA_DB.Save(&cluster).Error
}

func (clusterService *ClusterService) DeleteCluster(cluster system.SysCluster) (err error) {
	var entity system.SysCluster
	err = global.GVA_DB.Where("id = ?", cluster.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (clusterService *ClusterService) GetClusterInfoList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysCluster{})
	var clusterList []system.SysCluster

	err = db.Count(&total).Error

	if err != nil {
		return clusterList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["ip_group"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return clusterList, total, err
			}

			err = db.Order(OrderStr).Find(&clusterList).Error
		} else {
			err = db.Order("ip_group").Find(&clusterList).Error
		}
	}
	return clusterList, total, err
}

func (clusterService *ClusterService) GetAllClusters() (clusters []system.SysCluster, err error) {
	err = global.GVA_DB.Find(&clusters).Error
	return
}

func (clusterService *ClusterService) GetClusterById(id int) (cluster system.SysCluster, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&cluster).Error
	return
}
