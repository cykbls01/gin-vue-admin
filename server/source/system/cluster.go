package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initCluster struct{}

const initOrderCluster = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderCluster, &initCluster{})
}

func (i initCluster) InitializerName() string {
	return sysModel.SysCluster{}.TableName()
}

func (i *initCluster) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysCluster{})
}

func (i *initCluster) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysCluster{})
}

func (i *initCluster) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysCluster{
		{IpGroup: "43.128.79.149", Name: "腾讯云测试集群", Config: "uploads/file/43.128.79.149", Description: "腾讯云测试集群"},
		{IpGroup: "10.30.32.1", Name: "local测试集群", Config: "uploads/file/10.30.32.1", Description: "local测试集群"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysCluster{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initCluster) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "43.128.79.149").
		First(&sysModel.SysCluster{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
