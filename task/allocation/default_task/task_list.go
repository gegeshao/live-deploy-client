package default_task

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "log"
  "reflect"
  "strings"

  "github.com/yuin/gopher-lua"
)

type DefaultTaskMaid struct{}
func (maid *DefaultTaskMaid) UpdateScripts(task *schema.Task, L *lua.LState)(bool, string){
  filename:=task.Action
  url:=task.Content
  //保证安全
  filename = strings.Replace(filename, "/", "_", -1)
  if err:= utils.Download(url, "scripts/"+filename); err!=nil{
    log.Println(err)
    return false, err.Error()
  }
  L.DoString()
  return true, filename
}


func DoDefalutTask(task *schema.Task, L *lua.LState) (exist bool, status bool,result string) {
  taskType := task.Type
  config:=utils.GetConfig()
  allow := false
  for _, value := range config.LoadDefaultTask{
    if value == taskType{
      allow = true
    }
  }
  if !allow{
    return
  }
  defaultTask := &DefaultTaskMaid{}
  method:= reflect.ValueOf(defaultTask).MethodByName(taskType)
  if method.Kind().String() == "invalid"{
    return
  }
  returnValue := method.Call([]reflect.Value{reflect.ValueOf(task), reflect.ValueOf(L)})
  return true, returnValue[0].Bool(), returnValue[1].String()
}



