package lua_module

/*
* Get for https://github.com/kohkimakimoto/gluayaml/blob/master/gluayaml.go
*/
import (
  "live-deploy-client/utils"
  "github.com/yuin/gopher-lua"
)

func fromInterface(L *lua.LState, value interface{}) lua.LValue {
  switch converted := value.(type) {
  case bool:
    return lua.LBool(converted)
  case float64:
    return lua.LNumber(converted)
  case int:
    return lua.LNumber(converted)
  case string:
    return lua.LString(converted)
  case []interface{}:
    arr := L.CreateTable(len(converted), 0)
    for _, item := range converted {
      arr.Append(fromInterface(L, item))
    }
    return arr
  case map[interface{}]interface{}:
    tbl := L.CreateTable(0, len(converted))
    for key, item := range converted {
      if s, ok := key.(string); ok {
        tbl.RawSetH(lua.LString(s), fromInterface(L, item))
      }
    }
    return tbl
  }

  return lua.LNil
}

func GetConfig(L *lua.LState) int{
  config := utils.GetConfig().Plugin
  L.Push(fromInterface(L, config))
  return 1
}
