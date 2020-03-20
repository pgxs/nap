package main

import (
	"fmt"
	"openpitrix.io/libconfd"
	_ "openpitrix.io/libconfd/etcdv3/backend"
	"os"
)

func main() {
	cfg := libconfd.MustLoadConfig("./configs/agent.toml")
	backendCfg := libconfd.MustLoadBackendConfig("./configs/agent-backend.toml")
	backendClient, err := libconfd.NewBackendClient(backendCfg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	libconfd.NewProcessor().Run(cfg, backendClient)
}
