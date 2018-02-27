package vm

import ("github.com/yuin/gopher-lua"
  luaModule "live-deploy-client/vm/lua_module"
)

func Loader(L *lua.LState) int{
  mod := L.SetFuncs(L.NewTable(), map[string]lua.LGFunction{
    "execute":              luaModule.Execute,
    "executeCWD": luaModule.ExecuteCWD,
    "getConfig": luaModule.GetConfig,
    "path": luaModule.Path,
  })
  L.Push(mod)
  return 1
}