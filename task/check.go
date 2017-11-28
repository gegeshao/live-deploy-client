package task

import (
  "net/http"
  "live-deploy-client/utils"
  "fmt"
  "time"
  "encoding/json"
  "bytes"
)

func Check() error{
  config:= utils.GetConfig()
  jsonBody := map[string]interface{}{
    "now": time.Now().Unix(),
    "id": config.MachineID,
  }
  sendData, _:= json.Marshal(jsonBody)

  encrypData, err:= utils.CFBEncrypt([]byte(config.PrivateKey), sendData)
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