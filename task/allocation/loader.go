package allocation

import ("github.com/yuin/gopher-lua"
  luaModule "live-deploy-client/task/allocation/lua_module"
)

func Loader(L *lua.LState) int{
  mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
    "execute":              luaModule.Execute,
    "getConfig": luaModule.GetConfig,
    "path": luaModule.Path,
  })
  L.Push(mod)
  return 1
}