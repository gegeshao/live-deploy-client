package allocation

import (
  "live-deploy-client/schema"
  "live-deploy-client/utils"
  "live-deploy-client/vm"
  "log"
  "os"
  "testing"
)



func TestDoDefalutTask(t *testing.T) {
  finishTask := Dispatch(&schema.Task{
    Type: "UpdateScripts",
    Action:"nginx.lua",
    Content: "xxxxx.com/xxx.lua",
  })
  log.Println(finishTask)
}
