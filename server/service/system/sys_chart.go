package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type ChartService struct{}

func (chartService *ChartService) CreateChart(chart system.SysChart) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ?", chart.Name).First(&system.SysChart{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同chart")
	}
	return global.GVA_DB.Create(&chart).Error
}

func (chartService *ChartService) UpdateChart(chart system.SysChart) (err error) {
	return global.GVA_DB.Save(&chart).Error
}

func (chartService *ChartService) DeleteChart(chart system.SysChart) (err error) {
	var entity system.SysChart
	err = global.GVA_DB.Where("id = ?", chart.ID).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) { // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (chartService *ChartService) GetAllCharts() (charts []system.SysChart, err error) {
	err = global.GVA_DB.Find(&charts).Error
	return
}

func (chartService *ChartService) GetChartById(id int) (chart system.SysChart, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&chart).Error
	return
}

func (chartService *ChartService) GetChartInfoList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysChart{})
	var chartList []system.SysChart

	err = db.Count(&total).Error

	if err != nil {
		return chartList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["name"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return chartList, total, err
			}

			err = db.Order(OrderStr).Find(&chartList).Error
		} else {
			err = db.Order("name").Find(&chartList).Error
		}
	}
	return chartList, total, err
}
