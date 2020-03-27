package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"pgxs.io/chassis/log"

	"pgxs.io/nap/pkg/thirdparty/libconfd"
	_ "pgxs.io/nap/pkg/thirdparty/libconfd/etcdv3/backend"
)

type _TemplateResourceConfig struct {
	TemplateResource libconfd.TemplateResource `toml:"template"`
}

func main() {
	//
	res := libconfd.TemplateResource{Src: "test.toml", Dest: "test.conf", Keys: []string{"/test"},ReloadCmd:"docker ps",CheckCmd:"docker ps -a"}
	resCfg := &_TemplateResourceConfig{res}

	if err := toml.NewEncoder(os.Stdout).Encode(resCfg); err != nil {
		log.New().Fatal(err)
	}



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
