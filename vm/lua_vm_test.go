package vm

import (
  "live-deploy-client/utils"
  "log"
  "os"
  "testing"
)

var CONFIGPATH = "/Users/hyh/go-work/src/live-deploy-client/conf/config.yaml"

func TestMain(m *testing.M) {
  config, err:=utils.InitConfig(CONFIGPATH)
  if err!=nil{
    log.Println(err)
  }
  config.LuaScriptsDir = "/Users/hyh/go-work/src/live-deploy-client/scripts"
  if err:=InitLuaVM();err!=nil{
    log.Println(err)
  }
  InitLuaVM()
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}
func TestLvm(t *testing.T){
  L:= GetLuaVM()
  L.DoString(`
    gosystem = require('gosystem')

    print(gosystem.getConfig())
  `)
}
