package utils

import (
	"io/ioutil"

	"fmt"
	"gopkg.in/yaml.v2"
	"path"
  "net/url"
)



//Config 配置文件
type Config struct {
	Server string `yaml:"server"`
	NginxConfigPath string   `yaml:"nginx_config_path"`
  MachineID string `yaml:"machine_id"`
  PrivateKey string `yaml:"private_key"`
  CheckServer string `yaml:"-"`
  TaskServer string `yaml:"-"`
  LoadDefaultTask []string `yaml:"load_default_task"`
  LuaScriptsDir string `yaml:"-"`
}

var config *Config
var configInterface interface{}
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
	//校验配置
	if u, err:= url.Parse(config.Server); err!=nil{
	  fmt.Println("服务器地址配置错误")
	  return nil, err
  }else{
    yaml.Unmarshal(configBytes, &configInterface)
    u.Path = path.Join(u.Path, "/client/task")
    checkU, _:= url.Parse(config.Server)
    checkU.Path = path.Join(checkU.Path, "/client/check")
    config.TaskServer = u.String()
    config.CheckServer = checkU.String()
  }
  config.LuaScriptsDir = "scripts"


	return config, nil
}

//GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}

func GetConfigInterface() interface{} {
  return configInterface
}


func GetNginxTemplateFileName(id int64) string {
	return path.Join(config.NginxConfigPath, fmt.Sprintf("id-%d.conf", id))
}
