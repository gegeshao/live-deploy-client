package allocation

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "log"

  "github.com/yuin/gopher-lua"
)
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
  log.Println(L.Get(-3), L.Get(-1), L.Get(0), L.Get(1), L.Get(2))
}

func TaskFail(task *schema.Task, err error){
  log.Println(err)
}