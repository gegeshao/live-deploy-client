package allocation

import ("github.com/yuin/gopher-lua"
  luaModule "live-deploy-client/task/allocation/lua_module"
)

func Loader(L *lua.LState) int{
  mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
    "execute":              luaModule.Execute,
    "f": func(L *lua.LState) int{
      L.Push(lua.LString("hello"))
      return 1
    },
  })
  L.Push(mod)
  return 1
}