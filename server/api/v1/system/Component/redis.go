package Component

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"go.uber.org/zap"
	"os/exec"
)

type Redis struct {
}

func (redis *Redis) Create(str string) error {
	//str = str + fmt.Sprintf(" --set password=%s", component.Password)
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (redis *Redis) Delete(str string) error {
	cmd := exec.Command("/bin/bash", "-c", str)
	_, err := cmd.Output()
	if err != nil {
		global.GVA_LOG.Error(str, zap.Error(err))
	}
	return err
}

func (redis *Redis) Port(cluster system.SysCluster, component system.SysComponent) string {

	return fmt.Sprintf("kubectl get svc -n %s -l svc=nodePort -o jsonpath=\"{.items[0].spec.ports[0].nodePort}\" --kubeconfig %s", component.Namespace, global.GVA_CONFIG.Helm.Path+cluster.Config)
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
