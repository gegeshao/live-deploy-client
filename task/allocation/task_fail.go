package allocation

import (
  "log"
  "live-deploy-client/schema"
  "fmt"
)
func TaskFail(task *schema.Task, errMsg string)schema.TaskClientFinish{
  log.Println("task失败", errMsg)

  return schema.TaskClientFinish{
    ID: task.TaskID,
    Status: false,
    Content:errMsg,
  }
}