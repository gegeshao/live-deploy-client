package allocation

import (
  "fmt"
  "live-deploy-client/schema"
  "live-deploy-client/vm"

  "github.com/yuin/gopher-lua"
)

func DoCustomTask(task *schema.Task) schema.TaskClientFinish{

  task.TaskID = task.ID
  //对比数据库检查该任务是否已经完成
  if existTask, err:=schema.GetTaskByID(task.ID); existTask !=nil && err == nil{
    status:= false
    if existTask.Status == 1 {
      status = true
    }
    //任务已完成
    return schema.TaskClientFinish{
      ID: task.TaskID,
      Status: status,
      Result:  existTask.Result,
    }
  }

  L := vm.GetLuaVM()

  if L.GetGlobal(task.Type)==lua.LNil{
    return TaskFail(task, "脚本插件错误: 脚本缺少任务类型:"+task.Action)
  }
  tab:=L.GetGlobal(task.Type).(*lua.LTable)
  if L.GetField(tab, task.Action) == lua.LNil{
    return TaskFail(task, "脚本插件错误: 脚本缺少任务函数:"+task.Type)
  }
  p:=lua.P{
    Fn: L.GetField(tab, task.Action),
    NRet: 1,
    Protect:true,
  }
  if err:= L.CallByParam(p, lua.LNumber(task.TrackID), lua.LString(task.TrackKey), lua.LString(task.Content)); err!=nil{
    return TaskFail(task, fmt.Sprintf("%v", err))

  }
  ret := L.Get(-1)
  L.Pop(1)
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
  //记录完成任务到本地日志
  schema.AddTask(task)
  return schema.TaskClientFinish{
    ID: task.TaskID,
    Status: true,
    Result:  result,
  }
}

