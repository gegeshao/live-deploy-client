package task

import (
  "live-deploy-client/utils"
  "net/http"
  "log"
)
var (
  client = &http.Client{}
)
func Get(){
  config:=utils.GetConfig()
  machineKey := config.MachineID


  // 获取已完成任务列表

  req, _ := http.NewRequest("POST", config.Server, nil)
  req.Header.Set("private-key", machineKey)
  resp, err:=client.Do(req)
  if err!=nil{
    log.Println("获取任务失败", err)
  }

  if resp.StatusCode != 200{
    log.Println("获取任务失败: ", resp.StatusCode)
  }
  //没有任务
  if resp.Header.Get("task-count") == "none"{
    return
  }


}
