package allocation

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"

  "github.com/yuin/gopher-lua"
)

func DoTask(task *schema.Task) schema.TaskClientFinish{
  //TODO 检查是否已经完成
  L:=lua.NewState()
  defer L.Close()
  L.PreloadModule("gosystem", Loader)
  script, err:=utils.GetScript(task.Type)
  if err!=nil{
    return TaskFail(task,"没有任务处理模板")

  }
  if err:= L.DoString(script); err!=nil{
    return TaskFail(task, fmt.Sprintf("%v", err))

  }
  fn:=L.GetGlobal(task.Type).(*lua.LTable)

  p:=lua.P{
    Fn: L.GetField(fn, task.Action),
    NRet: 1,
    Protect:true,
  }
  if err:= L.CallByParam(p, lua.LNumber(task.TrackID), lua.LString(task.TrackKey), lua.LString(task.Content)); err!=nil{
    return TaskFail(task, fmt.Sprintf("%v", err))

  }
  ret := L.Get(-1)
  tabl, ok:= ret.(*lua.LTable)
  if !ok{
    return TaskFail(task, "脚本插件错误: 错误返回值")

  }

  status:=tabl.RawGetString("status")
  scriptResult:= tabl.RawGetString("result")
  success := false
  if status.Type() != lua.LTBool {
    return TaskFail(task, "脚本插件错误: 返回值字段类型错误, status 必须为bool")

  }
  success = lua.LVAsBool(status)
  if scriptResult.Type() != lua.LTString{
    return TaskFail(task, "脚本插件错误: 返回值字段类型错误, result 必须为 string")

  }
  result:= lua.LVAsString(scriptResult)
  if !success {
    return TaskFail(task, result)

  }
  //log.Println("任务成功:", result)
  return schema.TaskClientFinish{
    ID: task.TaskID,
    Status: true,
    Content:  result,
  }
}

