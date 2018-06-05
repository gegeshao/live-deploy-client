package utils

import (
	"io/ioutil"
  "os"
  "path/filepath"

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
  ProjectDir string `yaml:"project_dir"`
  LoadDefaultTask []string `yaml:"load_default_task"`
}

//Config 配置文件
type Config struct {
	System *System `yaml:"system"`
	Plugin interface{} `yaml:"plugin"`
  PrepareDir []string `yaml:"prepare_dir"`
  LuaScriptsDir string `yaml:"-"`
  CWD string `yaml:"-"`
}

var config *Config

func SetCWD(cwd string) error{
  if cwd == ""{
    ex, err := os.Executable()
    if err != nil {
      return err
    }
    cwd = filepath.Dir(ex)
  }
  if(config == nil){
    config = new(Config)
  }
  config.CWD = cwd
  return nil
}

func init(){
  SetCWD("")
}


func getProjectFilePath(filepath string) string{
  if path.IsAbs(filepath) {
    return filepath
  }
  return path.Join(config.CWD, filepath)
}

func initPrepareDir(){
  for _, filepath := range config.PrepareDir{
    os.MkdirAll(filepath, 0644)
  }
}

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

func initLuaScript() error{
  if config.LuaScriptsDir == ""{
    config.LuaScriptsDir = "scripts"
  }
  config.LuaScriptsDir = getProjectFilePath(config.LuaScriptsDir)

  return os.MkdirAll(config.LuaScriptsDir, 0644)
}

func initProjectDir()error{
  if config.System.ProjectDir == ""{
    config.System.ProjectDir = "projects"
  }
  if !path.IsAbs(config.System.ProjectDir){
    config.System.ProjectDir = getProjectFilePath(config.System.ProjectDir)
  }
  return os.MkdirAll(config.System.ProjectDir, 0644)
}


//InitConfig 读取配置文件
func InitConfig(source string) (*Config, error) {
	configBytes, err := ioutil.ReadFile(getProjectFilePath(source))
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
  if err := initLuaScript();err!=nil{
    return nil, err
  }
  if err := initProjectDir();err!=nil{
    return nil, err
  }
  initPrepareDir()
	return config, nil
}

//GetConfig 获取配置文件
func GetConfig() *Config {
	return config
}