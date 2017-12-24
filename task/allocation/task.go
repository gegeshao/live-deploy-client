package allocation

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"

  "github.com/yuin/gopher-lua"
  "log"
)

func DoTask(task *schema.Task){
  L:=lua.NewState()
  defer L.Close()
  L.PreloadModule("gosystem", Loader)
  script, err:=utils.GetScript(task.Type)
  if err!=nil{
    TaskFail(task, fmt.Errorf("没有任务处理模板"))
    return
  }
  if err:= L.DoString(script); err!=nil{
    TaskFail(task, err)
    return
  }
  fn:=L.GetGlobal(task.Type).(*lua.LTable)

  p:=lua.P{
    Fn: L.GetField(fn, task.Action),
    NRet: 1,
    Protect:true,
  }
  if err:= L.CallByParam(p, lua.LNumber(task.TrackID), lua.LString(task.TrackKey), lua.LString(task.Content)); err!=nil{
    TaskFail(task,err)
    return
  }
  ret := L.Get(-1)
  tabl, ok:= ret.(*lua.LTable)
  if !ok{
    TaskFail(task, fmt.Errorf("脚本插件错误: 错误返回值"))
    return
  }

  status:=tabl.RawGetString("status")
  scriptResult:= tabl.RawGetString("result")
  success := false
  if status.Type() != lua.LTBool {
    TaskFail(task, fmt.Errorf("脚本插件错误: 返回值字段类型错误, status 必须为bool"))
    return
  }
  success = lua.LVAsBool(status)
  if scriptResult.Type() != lua.LTString{
    TaskFail(task, fmt.Errorf("脚本插件错误: 返回值字段类型错误, result 必须为 string"))
    return
  }
  result:= lua.LVAsString(scriptResult)
  if !success {
    TaskFail(task, fmt.Errorf(result))
    return
  }
  log.Println("任务成功:", result)
}

