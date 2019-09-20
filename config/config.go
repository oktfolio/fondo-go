package config

import config "fondo-go/db"

var Conf *Config

// Config ...
type Config struct {
	EnvMode string `yaml:"env"`
	Name    string `yaml:"name"`
	Port    string `yaml:"port"`
	Envs    []*env `yaml:"envs"`
	Env     *env   `yaml:"-"`
}

type env struct {
	Name       string                     `yaml:"name"`
	Token      *token                     `yaml:"token"`
	Datasource []*config.DatasourceConfig `yaml:"datasource"`
}

type token struct {
	Issuer    string `yaml:"issuer"`
	ExpireIn  int64  `yaml:"expire_in"`
	SecretKey string `yaml:"secret_key"`
}

// InitConfig 初始化本项目配置
func InitConfig(configPath string) {
	Conf = &Config{

	}
	if err := config.InitConf(configPath, Conf); err != nil {
		panic(err)
	}
	for _, e := range Conf.Envs {
		if e.Name == Conf.EnvMode {
			Conf.Env = e
			config.InitDbsByConfig(Conf.Env.Datasource)
			break
		}
	}
}
