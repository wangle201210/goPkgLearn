package main

import (
	"context"

	"github.com/go-kratos/kratos/pkg/naming"
	"github.com/go-kratos/kratos/pkg/naming/discovery"
)

func reg() context.CancelFunc {
	ip := "127.0.0.1"
	port := "9009"
	region := "regionName"
	zone := "zoneName"
	host := "hostName"
	DeployEnv := "DeployEnvName"

	dis := discovery.New(&discovery.Config{
		Region: region,
		Zone: zone,
		Host: host,
		Env: DeployEnv,
		Nodes:[]string{"127.0.0.1:7171"},// 这个很重要
	})
	ins := &naming.Instance{
		Zone:     "zoneName",
		Env:      "DeployEnvName",
		AppID:    "wanna.app1", // 唯一
		Hostname: "hostName",
		Addrs: []string{
			"grpc://" + ip + ":" + port,
		},
	}

	cancel, err := dis.Register(context.Background(), ins)
	if err != nil {
		panic(err)
	}

	return cancel
}


func main() {
	cancel := reg()
	for  {

	}
	defer cancel()
}
