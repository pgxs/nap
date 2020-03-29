// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package libconfd

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"text/template"
)

type Config struct {
	// The path to confd configs.
	// If the confdir is rel path, must convert to abs path.
	//
	// abspath = filepath.Join(ConfigPath, Config.ConfDir)
	//
	ConfDir string `yaml:"confdir"`

	// The backend polling interval in seconds. (10)
	Interval int `yaml:"interval"`

	// Enable noop mode. Process all template resources; skip target update.
	Noop bool `yaml:"noop"`

	// The string to prefix to keys. ("/")
	Prefix string `yaml:"prefix"`

	// sync without check_cmd and reload_cmd.
	SyncOnly bool `yaml:"sync-only"`

	// run once and exit
	Onetime bool `yaml:"onetime"`

	// enable watch support
	Watch bool

	// keep staged files
	KeepStageFile bool `yaml:"keep-stage-file"`

	// PGP secret keyring (for use with crypt functions)
	PGPPrivateKey string `toml:"pgp_private_key" json:"pgp_private_key" yaml:"pgp-private-key"`

	// ----------------------------------------------------

	FuncMap        template.FuncMap                               `toml:"-" json:"-" yaml:"-"`
	FuncMapUpdater func(m template.FuncMap, basefn *TemplateFunc) `toml:"-" json:"-" yaml:"-"`

	HookAbsKeyAdjuster  func(absKey string) (realKey string) `toml:"-" json:"-" yaml:"-"`
	HookOnCheckCmdDone  func(trName, cmd string, err error)  `toml:"-" json:"-" yaml:"-"`
	HookOnReloadCmdDone func(trName, cmd string, err error)  `toml:"-" json:"-" yaml:"-"`
	HookOnUpdateDone    func(trName string, err error)       `toml:"-" json:"-" yaml:"-"`
}

const defaultConfigContent = `
confdir: "./conf"
interval: 10
# Enable noop mode. Process all template resources; skip target update.
noop: false
# The string to prefix to keys. ("/")
prefix: "/"
# sync without check_cmd and reload_cmd.
sync-only: true
# level which confd should log messages ("DEBUG")
log-level: "DEBUG"
# run once and exit
onetime: true
# enable watch support
watch: false
# keep staged files
keep-stage-file: false
# PGP secret keyring (for use with crypt functions)
pgp-private-key:  ""
`

func newDefaultConfig() (p *Config) {
	p = new(Config)
	err := yaml.Unmarshal([]byte(defaultConfigContent), p)
	if err != nil {
		GetLogger().Panic(err)
	}
	if !filepath.IsAbs(p.ConfDir) {
		absdir, err := filepath.Abs(".")
		if err != nil {
			GetLogger().Panic(err)
		}
		p.ConfDir = filepath.Clean(filepath.Join(absdir, p.ConfDir))
	}
	return
}

//func MustLoadConfig(path string) *Config {
//	p, err := LoadConfig(path)
//	if err != nil {
//		GetLogger().Fatal(err)
//	}
//	return p
//}

//MustLoadConfig load yaml config
func MustLoadConfig(path string) *Config {
	p, err := LoadConfig(path)
	if err != nil {
		GetLogger().Fatal(err)
	}
	return p
}

//LoadConfig load config from yaml
func LoadConfig(filename string) (p *Config, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	p = new(Config)
	if err := yaml.NewDecoder(f).Decode(p); err == nil {
		if !filepath.IsAbs(p.ConfDir) {
			if absdir, err := filepath.Abs(filepath.Dir(filename)); err == nil {
				p.ConfDir = filepath.Clean(filepath.Join(absdir, p.ConfDir))
			}
		}
		return p, nil
	}
	return nil, err
}

func (p *Config) Valid() error {
	if !filepath.IsAbs(p.ConfDir) {
		return fmt.Errorf("ConfDir is not abs path: %s", p.ConfDir)
	}

	if !dirExists(p.ConfDir) {
		return fmt.Errorf("ConfDir not exists: %s", p.ConfDir)
	}

	if p.Interval < 0 {
		return fmt.Errorf("invalid Interval: %d", p.Interval)
	}
	//if p.LogLevel != "" && !newLogLevel(p.LogLevel).Valid() {
	//	return fmt.Errorf("invalid LogLevel: %s", p.LogLevel)
	//}

	return nil
}

func (p *Config) Save(name string) error {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(p); err != nil {
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(buf.String()); err != nil {
		return err
	}
	return nil
}

func (p *Config) Clone() *Config {
	q := *p

	// clone map
	if p.FuncMap != nil {
		q.FuncMap = make(template.FuncMap)
		for k, v := range p.FuncMap {
			q.FuncMap[k] = v
		}
	}

	return &q
}

func (p *Config) GetConfigDir() string {
	return filepath.Join(p.ConfDir, "conf.d")
}

func (p *Config) GetTemplateDir() string {
	return filepath.Join(p.ConfDir, "templates")
}

func (p *Config) GetDefaultTemplateOutputDir() string {
	return filepath.Join(p.ConfDir, "templates_output")
}
