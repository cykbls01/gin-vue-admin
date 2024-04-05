package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
)

type ComponentService struct{}

func (componentService *ComponentService) CreateComponent(component *system.SysComponent) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? AND namespace = ? And status != '已回收'", component.Name, component.Namespace).First(&system.SysComponent{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同component")
	}
	return global.GVA_DB.Create(component).Error
}

func (componentService *ComponentService) UpdateComponent(component system.SysComponent) (err error) {
	return global.GVA_DB.Save(&component).Error
}

func (componentService *ComponentService) DeleteComponent(component system.SysComponent) (err error) {
	var entity system.SysComponent
	err = global.GVA_DB.Where("id = ?", component.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                            // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (componentService *ComponentService) GetAllComponents() (components []system.SysComponent, err error) {
	err = global.GVA_DB.Find(&components).Error
	return
}

func (componentService *ComponentService) GetComponentById(id int) (component system.SysComponent, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&component).Error
	return
}

func (componentService *ComponentService) GetComponentInfoList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysComponent{})
	var componentList []system.SysComponent

	err = db.Count(&total).Error

	if err != nil {
		return componentList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["name"] = true
			orderMap["namespace"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return componentList, total, err
			}

			err = db.Order(OrderStr).Find(&componentList).Error
		} else {
			err = db.Order("name").Find(&componentList).Error
		}
	}
	return componentList, total, err
}
