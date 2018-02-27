package utils

import (
	"io/ioutil"

  "fmt"
	"gopkg.in/yaml.v2"
	"path"
  "net/url"
)

type System struct{
  Server string `yaml:"server"`
  MachineID string `yaml:"machine_id"`
  PrivateKey string `yaml:"private_key"`
  CheckServer string `yaml:"-"`
  TaskServer string `yaml:"-"`
  LoadDefaultTask []string `yaml:"load_default_task"`
}

//Config 配置文件
type Config struct {
	System *System `yaml:"system"`
	Plugin interface{} `yaml:"plugin"`
  LuaScriptsDir string `yaml:"-"`
}

var config *Config

func initServerConfig() error{
  //校验服务器配置
  if u, err:= url.Parse(config.System.Server); err!=nil{
    fmt.Println("服务器地址配置错误")
    return err
  }else{
    u.Path = path.Join(u.Path, "/client/task")
    checkU, _:= url.Parse(config.System.Server)
    checkU.Path = path.Join(checkU.Path, "/client/check")
    config.System.TaskServer = u.String()
    config.System.CheckServer = checkU.String()
  }
  return nil
}

func initLuaScript(){
  config.LuaScriptsDir = "scripts"
}


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
  if err := initServerConfig(); err!=nil{
    return nil, err
  }
  initLuaScript()
	return config, nil
}

//GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}