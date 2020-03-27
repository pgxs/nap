package main

import (
	"fmt"
	"os"

	"pgxs.io/nap/pkg/thirdparty/libconfd"
	_ "pgxs.io/nap/pkg/thirdparty/libconfd/etcdv3/backend"
)

func main() {
	cfg := libconfd.MustLoadConfig("./configs/agent.yml")
	backendCfg := libconfd.MustLoadBackendConfig("./configs/agent-backend.yml")
	//logger := log.New().Component("agent").Category("starter")
	//logger.Infof("%+v", backendCfg)
	fmt.Printf("%+v", backendCfg)
	backendClient, err := libconfd.NewBackendClient(backendCfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	libconfd.NewProcessor().Run(cfg, backendClient)
}
