package default_task

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
  "os"
  "testing"
)

var CONFIGPATH = "/Users/hyh/go-work/src/live-deploy-client/conf/config.yaml"

func TestMain(m *testing.M) {
  utils.InitConfig(CONFIGPATH)
  vm.InitLuaVM()
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}

func TestDoDefalutTask(t *testing.T) {
  exits, status, result:=DoDefalutTask(&schema.Task{
    Type: "UpdateScripts",
    Action:"nginx.lua",
    Content: "xxxxx.com/xxx.lua",
  })
  if !exits{
    t.Fail()
  }
  log.Println(status, result)
}
