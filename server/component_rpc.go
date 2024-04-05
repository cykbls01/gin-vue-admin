package main

import (
	"flag"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/rpc"
	"github.com/flipped-aurora/gin-vue-admin/server/core"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/smallnest/rpcx/server"
	"go.uber.org/zap"
)

func main() {
	global.GVA_VP = core.Viper() // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)

	flag.Parse()
	s := server.NewServer()
	s.RegisterName("HelmRpc", new(rpc.HelmRpc), "")
	err := s.Serve("tcp", *flag.String("addr1", "0.0.0.0:8972", "server address"))
	if err != nil {
		panic(err)
	}
}
