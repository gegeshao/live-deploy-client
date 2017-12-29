package default_task

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
  "reflect"
  "strings"
)

type DefaultTaskMaid struct{}
func (maid *DefaultTaskMaid) UpdateScripts(task *schema.Task)(bool, string){
  filename:=task.Action
  url:=task.Content
  //保证安全
  filename = strings.Replace(filename, "/", "_", -1)
  if err:= utils.Download(url, "scripts/"+filename); err!=nil{
    log.Println(err)
    return false, err.Error()
  }
  //加载新的脚本
  if err := vm.LoadScript(filename); err!=nil{
    return false, fmt.Sprintf("加载%s脚本失败: %v", filename, err)
  }
  return true, filename
}


func DoDefalutTask(task *schema.Task) (exist bool, status bool,result string) {
  taskType := task.Type
  config:=utils.GetConfig()
  allow := false
  for _, value := range config.System.LoadDefaultTask{
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
  returnValue := method.Call([]reflect.Value{reflect.ValueOf(task)})
  return true, returnValue[0].Bool(), returnValue[1].String()
}



