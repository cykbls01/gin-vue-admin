package Component

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"os/exec"
)

type Zookeeper struct {
}

func (Zookeeper *Zookeeper) Create(str string) error {
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (Zookeeper *Zookeeper) Delete(str string) error {
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (Zookeeper *Zookeeper) Port(cluster system.SysCluster, component system.SysComponent) string {
	return fmt.Sprintf("kubectl get --namespace %s -o jsonpath=\"{.spec.ports[0].nodePort}\" services %s --kubeconfig %s", component.Namespace, component.Name, global.GVA_CONFIG.Helm.Path+cluster.Config)
	//cmd := exec.Command("/bin/bash", "-c", str)
	//port, err := cmd.Output()
	//if err != nil {
	//	global.GVA_LOG.Error(str, zap.Error(err))
	//	return 0, err
	//}
	//
	//var p int
	//err = json.Unmarshal(port, &p)
	//if err != nil {
	//	global.GVA_LOG.Error(str, zap.Error(err))
	//	return 0, err
	//}
	//
	//return p, nil
}
