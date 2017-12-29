package vm

import (
  "io/ioutil"
  "live-deploy-client/utils"
  "github.com/yuin/gopher-lua"
)

func getScript(filename string)(string, error){
  fileContent, err:=ioutil.ReadFile("scripts/"+filename)
  if err!=nil{
    return "", nil
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
    fileContent, err:= getScript(file.Name())
    if err!=nil{
      return err
    }
    if err := L.DoString(fileContent); err!=nil{
      return  err
    }
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


