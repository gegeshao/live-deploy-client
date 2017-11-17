package nginx

import (
	"bytes"
	"fmt"
	"live-deploy-client/utils"
	"os/exec"
)

func EnsureNginxConfig() (outMsg string, err error) {
	config := utils.GetConfig()
	cmd := exec.Command(config.NginxTest[0], config.NginxTest[1:]...)
	var outBuf bytes.Buffer
	cmd.Stderr = &outBuf
	startErr := cmd.Start()
	waritErr := cmd.Wait()
	out := outBuf.String()
	if startErr != nil || waritErr != nil {
		err = fmt.Errorf(out)
	} else {
		outMsg = out
	}
	return
}

func RestartNginx() (outMsg string, err error) {
	config := utils.GetConfig()
	cmd := exec.Command(config.NginxReload[0], config.NginxReload[1:]...)
	var outBuf bytes.Buffer
	cmd.Stderr = &outBuf
	startErr := cmd.Start()
	waritErr := cmd.Wait()
	out := outBuf.String()
	if startErr != nil || waritErr != nil {
		err = fmt.Errorf(out)
	} else {
		outMsg = out
	}
	return
}
