package task

import (
  "net/http"
  "live-deploy-client/utils"
  "fmt"
  "time"
  "encoding/json"
  "bytes"
  "github.com/huyinghuan/cfb"
)

/*
时间戳 + 机器id 用 private_key 进行加密发送
*/
func Check() error{
  config:= utils.GetConfig()
  jsonBody := map[string]interface{}{
    "now": time.Now().Unix(),
    "id": config.MachineID,
  }
  sendData, _:= json.Marshal(jsonBody)

  encrypData, err:= cfb.Encrypt([]byte(config.PrivateKey), sendData)
  if err!=nil{
    return err
  }

  req, err:= http.NewRequest("POST", config.CheckServer , bytes.NewBuffer(encrypData))
  if err!=nil{
    return err
  }
  req.Header.Set("private-key", config.MachineID)
  resp, err:= client.Do(req)
  if err!=nil{
    return err
  }
  if resp.StatusCode != 200{
    return fmt.Errorf("验证错误, status : %d", resp.StatusCode)
  }
  return nil
}

//检查任务是否已完成
func CheckTaskIsFinish() bool{
  return false
}