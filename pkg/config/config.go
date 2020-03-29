package config

import (
	cfg "pgxs.io/chassis/config"
)

//ServerConfig
type ServerConfig struct {
}

type EtcdConfig struct {
	Endpoints []string `yaml: ,flow`
	Host      []string `yaml:",flow"`

	UserName string `yaml:"user"`
	Password string `yaml:"password"`

	ClientCAKeys string `yaml:"client-ca-keys"`
	ClientCert   string `yaml:"client-cert"`
	ClientKey    string `yaml:"client-key"`
}

func Init() {
	//reset conf file env key
	cfg.SetLoadFileEnvKey("NAP_CONF")
	cfg.LoadFromEnvFile()
}
