package Component

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type ComponentTask interface {
	Create(str string) error
	Delete(str string) error
	Port(cluster system.SysCluster, component system.SysComponent) string
}

func Convert(ty string) ComponentTask {
	switch ty {
	case "redis":
		return new(Redis)
	case "rabbitmq":
		return new(Rabbitmq)
	case "zookeeper":
		return new(Zookeeper)
	}
	return nil
}
