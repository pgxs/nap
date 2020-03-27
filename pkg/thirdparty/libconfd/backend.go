// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package libconfd

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

type BackendConfig struct {
	Type string
	Host []string `yaml:",flow"`

	UserName string `yaml:"user"`
	Password string `yaml:"password"`

	ClientCAKeys string `yaml:"client-ca-keys"`
	ClientCert   string `yaml:"client-cert"`
	ClientKey    string `yaml:"client-key"`

	HookKeyAdjuster func(key string) (realKey string) `yaml:"-" json:"-"`
}

func (p *BackendConfig) Clone() *BackendConfig {
	var q = *p
	q.Host = append([]string{}, p.Host...)
	return &q
}

type BackendClient interface {
	Type() string
	GetValues(keys []string) (map[string]string, error)
	WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error)
	WatchEnabled() bool
	Close() error
}

func MustNewBackendClient(cfg *BackendConfig, opts ...func(*BackendConfig)) BackendClient {
	p, err := NewBackendClient(cfg, opts...)
	if err != nil {
		GetLogger().Panic(err)
	}
	return p
}

func NewBackendClient(cfg *BackendConfig, opts ...func(*BackendConfig)) (BackendClient, error) {
	cfg = cfg.Clone()
	for _, fn := range opts {
		fn(cfg)
	}

	newClient := _BackendClientMap[cfg.Type]
	if newClient == nil {
		return nil, fmt.Errorf("libconfd: unknown backend type %q", cfg.Type)
	}

	return newClient(cfg)
}

func MustLoadBackendConfig(path string) *BackendConfig {
	p, err := LoadBackendConfig(path)
	if err != nil {
		GetLogger().Fatal(err)
	}
	return p
}

func LoadBackendConfig(path string) (p *BackendConfig, err error) {
	p = new(BackendConfig)
	//_, err = toml.DecodeFile(path, p)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func RegisterBackendClient(
	typeName string,
	newClient func(cfg *BackendConfig) (BackendClient, error),
) {
	_BackendClientMap[typeName] = newClient
}

var _BackendClientMap = map[string]func(cfg *BackendConfig) (BackendClient, error){}
