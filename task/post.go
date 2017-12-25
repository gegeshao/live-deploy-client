package task

import (
  "log"
  "net/http"
	"bytes"
	"fmt"
	"live-deploy-client/utils"
)

func Post(sendData []byte) error{
	config:= utils.GetConfig()

	req, err:= http.NewRequest("POST", utils.GetAPIPath("/client/task/done") , bytes.NewBuffer(sendData))
	if err!=nil{
		return err
	}
	req.Header.Set("private-key", config.MachineID)
	resp, err:= client.Do(req)
	if err!=nil{
    log.Println(err, 999999)
		return err
	}
	if resp.StatusCode == 200{
	  log.Println("提交数据成功")
	  return nil
  }
	if resp.StatusCode == 403{
		return fmt.Errorf("验证错误, status : %d", resp.StatusCode)
	}

	return fmt.Errorf("服务器错误, status : %d", resp.StatusCode)
}
