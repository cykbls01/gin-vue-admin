package Component

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"os/exec"
)

type Rabbitmq struct {
}

func (rabbitmq *Rabbitmq) Create(str string) error {
	//str = str + fmt.Sprintf(" --set auth.password=%s", component.Password)
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (rabbitmq *Rabbitmq) Delete(str string) error {
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (rabbitmq *Rabbitmq) Port(cluster system.SysCluster, component system.SysComponent) string {
	if component.Name != "rabbitmq" {
		component.Name = component.Name + "-rabbitmq"
	}
	return fmt.Sprintf("kubectl get --namespace %s -o jsonpath=\"{.spec.ports[?(@.name=='amqp')].nodePort}\" services %s --kubeconfig %s", component.Namespace, component.Name, global.GVA_CONFIG.Helm.Path+cluster.Config)
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
	//	return 0, err
	//}
	//
	//return p, nil
}
