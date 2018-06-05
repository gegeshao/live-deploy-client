package vm

import (
  "io/ioutil"
  "live-deploy-client/utils"
  "github.com/yuin/gopher-lua"
  "log"
  "path"
)

func getScript(filepath string)(string, error){
  fileContent, err:=ioutil.ReadFile(filepath)
  if err!=nil{
    return "", err
  }
  return string(fileContent), nil
}

type lvm struct{
  luaState *lua.LState
}

var vm = &lvm{}

func GetLuaVM() *lua.LState{
  return vm.luaState
}

func InitLuaVM()(error){
  config:=utils.GetConfig()
  if vm.luaState != nil{
    return nil
  }
  L := lua.NewState()
  L.PreloadModule("gosystem", Loader)
  fileList, err:= ioutil.ReadDir(config.LuaScriptsDir)
  if err!= nil{
    return err
  }
  for _, file := range fileList{
    if file.IsDir(){
      continue
    }
    fileContent, err:= getScript(path.Join(config.LuaScriptsDir, file.Name()))
    if err!=nil{
      return err
    }
    if err := L.DoString(fileContent); err!=nil{
      return  err
    }
    log.Printf("加载 %s 插件成功!\n", file.Name())
  }
  vm.luaState = L
  return nil
}

func LoadScript(filename string) (error){
  if content, err:= getScript(filename); err!=nil{
    return err
  }else{
    vm.luaState.DoString(content)
    return nil
  }
}


