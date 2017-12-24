package task

import (
	"net/http"
	"bytes"
	"fmt"
	"live-deploy-client/utils"
)

func Post(sendData []byte) error{
	config:= utils.GetConfig()

	req, err:= http.NewRequest("POST", utils.GetAPIPath("/client/check") , bytes.NewBuffer(sendData))
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
