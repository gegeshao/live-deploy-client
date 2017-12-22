package allocation

import (
  "log"
  "live-deploy-client/schema"
)
func TaskFail(task *schema.Task, err error){
  log.Println("task失败",err)
}