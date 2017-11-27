package utils

import (
	"io/ioutil"

	"fmt"
	"gopkg.in/yaml.v2"
	"path"
)



//Config 配置文件
type Config struct {
	Dev bool `yaml:"dev"`
	Server string `yaml:"server"`
	NginxTest       []string `yaml:"nginx_test"`
	NginxReload     []string `yaml:"nginx_reload"`
	NginxConfigPath string   `yaml:"nginx_config_path"`
  MachineID string `yaml:"machine_id"`
}

var config *Config

//InitConfig 读取配置文件
func InitConfig(source string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(source)
	if err != nil {
		return nil,err
	}
	err = yaml.Unmarshal(configBytes, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

//GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}


func GetNginxTemplateFileName(id int64) string {
	return path.Join(config.NginxConfigPath, fmt.Sprintf("id-%d.conf", id))
}
