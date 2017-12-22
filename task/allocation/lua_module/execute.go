package lua_module

import (
  "bytes"
  "os/exec"
  "github.com/yuin/gopher-lua"
)

func Execute(L *lua.LState) int{
  params:= []string{}
  command := L.CheckString(1)
  for i:=2; i <= L.GetTop(); i++{
    params = append(params, L.CheckString(i))
  }
  cmd := exec.Command(command, params...)
  var outBuf bytes.Buffer
  cmd.Stderr = &outBuf
  startErr := cmd.Start()
  waritErr := cmd.Wait()
  out := outBuf.String()
  if startErr != nil || waritErr != nil {
    L.Push(lua.LBool(false))
    L.Push(lua.LString(out))
  } else {

    L.Push(lua.LBool(true))
    L.Push(lua.LString(out))
  }
  return 2
}