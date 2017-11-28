package task

import (
  "live-deploy-client/utils"
  "fmt"
  "net/http"
)
var (
  client = &http.Client{}
)
func Get(){
  config:=utils.GetConfig()
  machineKey := config.MachineID
  fmt.Println(machineKey)
  // 获取已完成任务列表

  req, _ := http.NewRequest("POST", config.Server, nil)
  client.Do(req)

}
