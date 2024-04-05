package system

import (
	"context"
	sysModel "github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type initChart struct{}

const initOrderChart = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderChart, &initChart{})
}

func (i initChart) InitializerName() string {
	return sysModel.SysChart{}.TableName()
}

func (i *initChart) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysChart{})
}

func (i *initChart) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysChart{})
}

func (i *initChart) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysChart{
		{Name: "redis", Path: "bitnami/redis-cluster", Config: "uploads/file/redis.yaml", Description: "redis集群", Category: "redis"},
		{Name: "rabbitmq", Path: "bitnami/rabbitmq", Config: "uploads/file/rabbitmq.yaml", Description: "rabbitmq集群", Category: "rabbitmq"},
		{Name: "zookeeper", Path: "bitnami/zookeeper", Config: "uploads/file/zookeeper.yaml", Description: "zookeeper集群", Category: "zookeeper"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, sysModel.SysChart{}.TableName()+"表数据初始化失败!")
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initChart) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "43.128.79.149").
		First(&sysModel.SysChart{}).Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}
