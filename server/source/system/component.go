package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initComponent struct{}

const initOrderComponent = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderComponent, &initComponent{})
}

func (i initComponent) InitializerName() string {
	return sysModel.SysComponent{}.TableName()
}

func (i *initComponent) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysComponent{})
}

func (i *initComponent) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysComponent{})
}

func (i *initComponent) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysComponent{
		{Name: "测试用"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysComponent{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initComponent) DataInserted(ctx context.Context) bool {
	return true
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "43.128.79.149").
		First(&sysModel.SysComponent{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return false
}
