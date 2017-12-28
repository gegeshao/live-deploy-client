package default_task

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "log"
  "os"
  "testing"
)

var CONFIGPATH = "/Users/hyh/go-work/src/live-deploy-client/conf/config.yaml"

func TestMain(m *testing.M) {
  utils.InitConfig(CONFIGPATH)
  // call flag.Parse() here if TestMain uses flags
  os.Exit(m.Run())
}

func TestDoDefalutTask(t *testing.T) {
  exits, status, result:=DoDefalutTask(&schema.Task{
    Type: "Download",
    Action:"nginx.lua",
    Content: "xxxxx.com/xxx.lua",
  }, nil)
  if !exits{
    t.Fail()
  }
  log.Println(status, result)
}
