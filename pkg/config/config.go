package config

import (
	cfg "pgxs.io/chassis/config"
)

type CustomConfig struct {
}

type EtcdConfig struct {
	Endpoints []string `yaml: ,flow`
}

func Init() {
	//reset conf file env key
	cfg.SetLoadFileEnvKey("NAP_CONF")
	cfg.LoadFromEnvFile()

}
