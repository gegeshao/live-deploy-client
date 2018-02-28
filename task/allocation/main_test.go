package allocation

import (
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "os"
  "testing"
)

var CWD = "/Users/hyh/go-work/src/live-deploy-client"

func TestMain(m *testing.M) {
  utils.SetCWD(CWD)
  utils.InitConfig("./conf/config.yaml")
  vm.InitLuaVM()
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}




