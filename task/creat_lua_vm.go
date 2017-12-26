package task

import (
  "io/ioutil"
  "live-deploy-client/task/allocation"
  "github.com/yuin/gopher-lua"
)

func getScript(filename string)(string, error){
  fileContet, err:=ioutil.ReadFile("scripts/"+filename)
  if err!=nil{
    return "", nil
  }
  return string(fileContet), nil
}


func CreateLuaVM()(*lua.LState, error){
  L:=lua.NewState()
  L.PreloadModule("gosystem", allocation.Loader)
  fileList, err:= ioutil.ReadDir("scripts")
  if err!= nil{
    return nil, err
  }
  for _, file := range fileList{
    if file.IsDir(){
      continue
    }
    fileContent, err:= getScript(file.Name())
    if err!=nil{
      return nil, err
    }
    if err:=L.DoString(fileContent); err!=nil{
      return nil, err
    }
  }

  return L, err
}
