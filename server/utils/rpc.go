package utils

import (
	"context"
	"flag"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

var Addr = flag.String("addr", "192.168.153.161:8972", "server address")

//var Addr = flag.String("addr", "localhost:8972", "server address")

func Call(path, method string, args interface{}) (error, system.SysComponent) {
	flag.Parse()

	d, _ := client.NewPeer2PeerDiscovery("tcp@"+*Addr, "")
	opt := client.DefaultOption
	opt.SerializeType = protocol.JSON

	xclient := client.NewXClient(path, client.Failtry, client.RandomSelect, d, opt)
	defer xclient.Close()

	reply := &system.SysComponent{}
	err := xclient.Call(context.Background(), method, args, reply)
	return err, *reply
}
