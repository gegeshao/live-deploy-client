package allocation

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
  "os"
  "testing"
)

var CWD = "/Users/hyh/go-work/src/live-deploy-client"

func TestMain(m *testing.M) {
  utils.SetCWD(CWD)
  utils.InitConfig("./conf/config.yaml")
  //初始化sqlite3
  if err:= schema.InitDriver(); err!=nil{
    log.Fatalln(err)
    return
  }
  vm.InitLuaVM()
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}
