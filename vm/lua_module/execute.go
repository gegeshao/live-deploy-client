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
  var errBuf bytes.Buffer
  var outBuf bytes.Buffer
  cmd.Stdout = &outBuf
  cmd.Stderr = &errBuf
  startErr := cmd.Start()
  waitErr := cmd.Wait()
  err := errBuf.String()
  out := outBuf.String()
  allOut:=""
  if err != ""{ allOut = allOut + err}
  if out != ""{allOut = allOut + out}
  if startErr != nil || waitErr != nil {
    L.Push(lua.LBool(false))
    L.Push(lua.LString(allOut))
  } else {
    L.Push(lua.LBool(true))
    L.Push(lua.LString(allOut))
  }
  return 2
}