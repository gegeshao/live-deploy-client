package task

import "live-deploy-client/utils"

func Get(){
  config:=utils.GetConfig()
  machineKey = config.Mac
}