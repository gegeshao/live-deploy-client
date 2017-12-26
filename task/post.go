package task

import (
  "log"
  "net/http"
	"bytes"
	"live-deploy-client/utils"
)

func Post(sendData []byte, err error){
  if err!=nil{
    log.Println(err)
    return
  }
	config:= utils.GetConfig()

	req, err:= http.NewRequest("POST", utils.GetAPIPath("/client/task/done") , bytes.NewBuffer(sendData))
	if err!=nil{
    log.Println(err)
		return
	}
	req.Header.Set("private-key", config.MachineID)
	resp, err:= client.Do(req)

	if err!=nil{
    log.Println(err)
		return
	}
	if resp.StatusCode == 200{
	  log.Println("提交数据成功")
	  return
  }
	if resp.StatusCode == 403{
		log.Printf("验证错误, status: %d\n", resp.StatusCode)
		return
	}

	log.Printf("服务器错误, status : %d\n", resp.StatusCode)
}
