package allocation

import (
  DefaultTask "live-deploy-client/task/allocation/default_task"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "reflect"
)


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
  defaultTask := &DefaultTask.DefaultTaskMaid{}
  method:= reflect.ValueOf(defaultTask).MethodByName(taskType)
  if method.Kind().String() == "invalid"{
    return
  }
  returnValue := method.Call([]reflect.Value{reflect.ValueOf(task)})
  return true, returnValue[0].Bool(), returnValue[1].String()
}