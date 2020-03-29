package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
	"pgxs.io/chassis/log"

	"pgxs.io/nap/pkg/thirdparty/libconfd"
)

const (
	confdBackendTypeEtcdv3 = "etcdv3"
)

var agentCfgLog = log.New().Component("config").Category("agent")

//AgentConfig nap agent config
type AgentConfig struct {
	Confd *libconfd.Config
	Etcd  *libconfd.BackendConfig
}

//MustLoadConfig load all config
func MustLoadAgentConfig(path string) *AgentConfig {
	aCfg, err := LoadAllConfig(path)
	if err != nil {
		agentCfgLog.Fatal(err)
	}
	return aCfg
}

//LoadAllConfig load confd & backend etcd config
func LoadAllConfig(path string) (cfg *AgentConfig, err error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	cfg = new(AgentConfig)
	if err := yaml.NewDecoder(f).Decode(cfg); err == nil {
		if !filepath.IsAbs(cfg.Confd.ConfDir) {
			if absDir, err := filepath.Abs(filepath.Dir(path)); err == nil {
				cfg.Confd.ConfDir = filepath.Clean(filepath.Join(absDir, cfg.Confd.ConfDir))
			}
		}
		cfg.Etcd.Type = confdBackendTypeEtcdv3
		return cfg, nil
	}
	return nil, err
}
