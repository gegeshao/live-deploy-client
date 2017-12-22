package lua_module

import (
  "path"

  "github.com/yuin/gopher-lua"
)

func join(L *lua.LState) int{
  if L.GetTop() == 1 {
    L.Push(L.Get(1))
    return 1
  }
  params:= []string{}
  for i:=1; i <= L.GetTop(); i++{
    params = append(params, L.CheckString(i))
  }
  L.Push(lua.LString(path.Join(params...)))
  return 1
}

func Path(L *lua.LState) int {
  mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
    "join": join,
  })
  L.Push(mod)
  return 1
}
