package allocation

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "log"

  "github.com/yuin/gopher-lua"
)

type LuaResult struct{
  Status string
  Result string
}
func DoTask(task *schema.Task){
  L:=lua.NewState()
  defer L.Close()
  content, err:= utils.GetScript(task.Type)
  if err!=nil{
    TaskFail(task, err)
    return
  }
  if err:=L.DoString(content); err!=nil{
    TaskFail(task, err)
    return
  }
  p:=lua.P{
    Fn: L.GetGlobal(task.Action),
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
  if status.Type() == lua.LTNumber{
    // 脚本退出状态码  1 错误退出  0 正常退出
    if lua.LVAsNumber(status) == 0 {
      success = true
    }
  }else if status.Type() == lua.LTBool{
    // 脚本返回码  true 正常 false 错误
    success = lua.LVAsBool(status)
  }else{
    TaskFail(task, fmt.Errorf("脚本插件错误: 返回值字段类型错误, status 必须为bool或0或1"))
    return
  }
  if scriptResult.Type() != lua.LTString{
    TaskFail(task, fmt.Errorf("脚本插件错误: 返回值字段类型错误, result 必须为 string"))
    return
  }
  result:= lua.LVAsString(scriptResult)
  if !success {
    TaskFail(task, fmt.Errorf(result))
    return
  }
}

func TaskFail(task *schema.Task, err error){
  log.Println(err)
}