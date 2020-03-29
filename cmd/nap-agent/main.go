package main

import (
	"fmt"
	"os"
	"pgxs.io/nap/pkg/config"

	"github.com/BurntSushi/toml"
	"pgxs.io/chassis/log"

	"pgxs.io/nap/pkg/thirdparty/libconfd"
	_ "pgxs.io/nap/pkg/thirdparty/libconfd/etcdv3/backend"
)

type _TemplateResourceConfig struct {
	TemplateResource libconfd.TemplateResource `toml:"template"`
}

func main() {
	config.Init()
	//libconfd.GetLogger().SetLevel("debug")
	//
	res := libconfd.TemplateResource{Src: "test.toml", Dest: "test.conf", Keys: []string{"/test"}, ReloadCmd: "docker ps", CheckCmd: "docker ps -a"}
	resCfg := &_TemplateResourceConfig{res}

	if err := toml.NewEncoder(os.Stdout).Encode(resCfg); err != nil {
		log.New().Fatal(err)
	}

	cfg := config.MustLoadAgentConfig("./configs/agent.yml")

	fmt.Printf("%+v\n%+v", cfg.Etcd, cfg.Confd)
	backendClient, err := libconfd.NewBackendClient(cfg.Etcd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	libconfd.NewProcessor().Run(cfg.Confd, backendClient)
}
