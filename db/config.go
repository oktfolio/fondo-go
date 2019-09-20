package db

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

// DBConfig 数据库配置
type DatasourceConfig struct {
	Host      string `yaml:"host"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	UseSSL    string `yaml:"use_ssl"`
	LogEnable bool   `yaml:"log"`
}

// InitConf 初始化配置
func InitConf(path string, conf interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, conf); err != nil {
		return err
	}
	return nil
}
