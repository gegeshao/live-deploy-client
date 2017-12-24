package task

import (
  "live-deploy-client/task/allocation"
  "live-deploy-client/utils"
  "net/http"
  "log"
  "io/ioutil"
  "encoding/gob"
  "bytes"
  "live-deploy-client/schema"
  "github.com/huyinghuan/cfb"
  "live-deploy-client/task"
)
var (
  client = &http.Client{}
)
func Get(){
  config:=utils.GetConfig()
  machineKey := config.MachineID
  cfbKey := config.PrivateKey
  // 获取已完成任务列表
  req, _ := http.NewRequest("POST", config.TaskServer, nil)
  req.Header.Set("private-key", machineKey)
  resp, err:=client.Do(req)
  if err!=nil{
    log.Println("获取任务失败", err)
    return
  }

  if resp.StatusCode != 200{
    log.Println("获取任务失败: ", resp.StatusCode)
    return
  }
  //没有任务
  if resp.Header.Get("task-count") == "none"{
    return
  }

  body, err:= ioutil.ReadAll(resp.Body)

  if err!=nil{
    log.Println("读取任务内容失败", err)
    return
  }

  r, err:= cfb.Decrypt([]byte(cfbKey), body)
  if err!=nil{
    log.Println("任务解析失败", err)
  }
  dec := gob.NewDecoder(bytes.NewReader(r))
  var taskList []schema.Task
  dec.Decode(&taskList)

  taskDoneList := []schema.TaskClientFinish{}
  for _, task:= range taskList{
    taskDone:=allocation.DoTask(&task)
    taskDoneList = append(taskDoneList, taskDone)
  }
  //发送完成状态
  resultBody, err:=utils.EncryptInterface(cfbKey, &taskDoneList)
  if err != nil{
    log.Println(err)
    return
  }
  task.Post()
}
